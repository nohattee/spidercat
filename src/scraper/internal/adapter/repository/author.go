package repository

import (
	"context"
	"time"

	"github.com/nohattee/spidercat/src/gopkg/ulid"
	"github.com/nohattee/spidercat/src/scraper/internal/domain/author"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

type authorModel struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (repo *AuthorRepository) GetOrCreateByNames(ctx context.Context, names []string) (author.Authors, error) {
	authorModels := make([]authorModel, len(names))
	for i := range names {
		authorModels[i] = authorModel{
			ID:   ulid.New(),
			Name: names[i],
		}
	}

	if err := repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "name"}},
			DoNothing: true,
		}).CreateInBatches(authorModels, 100)
		if result.Error != nil {
			return result.Error
		}

		result = tx.Where("name IN ?", names).Find(&authorModels)
		if result.Error != nil {
			return result.Error
		}
		return nil
	}); err != nil {
		return nil, err
	}

	categories := make(author.Authors, len(authorModels))
	for i := range categories {
		categories[i] = author.UnmarshalFromDB(authorModels[i].ID, authorModels[i].Name)
	}
	return categories, nil
}
