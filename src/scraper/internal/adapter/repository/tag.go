package repository

import (
	"context"
	"time"

	"github.com/nohattee/spidercat/src/gopkg/ulid"
	"github.com/nohattee/spidercat/src/scraper/internal/domain/tag"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{
		db: db,
	}
}

type tagModel struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (tagModel) TableName() string {
	return "tag"
}

func (repo *TagRepository) GetOrCreateByNames(ctx context.Context, names []string) (tag.Tags, error) {
	tagModels := make([]tagModel, len(names))
	for i := range names {
		tagModels[i] = tagModel{
			ID:        ulid.New(),
			Name:      names[i],
			UpdatedAt: time.Now(),
		}
	}

	if err := repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "name"}},
			DoNothing: true,
		}).CreateInBatches(tagModels, 100)
		if result.Error != nil {
			return result.Error
		}

		result = tx.Where("name IN ?", names).Find(&tagModels)
		if result.Error != nil {
			return result.Error
		}
		return nil
	}); err != nil {
		return nil, err
	}

	categories := make(tag.Tags, len(tagModels))
	for i := range categories {
		categories[i] = tag.UnmarshalFromDB(tagModels[i].ID, tagModels[i].Name)
	}
	return categories, nil
}
