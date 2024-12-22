package repository

import (
	"context"
	"time"

	"github.com/nohattee/spidercat/src/gopkg/ulid"
	"github.com/nohattee/spidercat/src/scraper/internal/domain/category"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

type categoryModel struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (repo *CategoryRepository) GetOrCreateByNames(ctx context.Context, names []string) (category.Categories, error) {
	categoryModels := make([]categoryModel, len(names))
	for i := range names {
		categoryModels[i] = categoryModel{
			ID:   ulid.New(),
			Name: names[i],
		}
	}

	if err := repo.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "name"}},
			DoNothing: true,
		}).CreateInBatches(categoryModels, 100)
		if result.Error != nil {
			return result.Error
		}

		result = tx.Where("name IN ?", names).Find(&categoryModels)
		if result.Error != nil {
			return result.Error
		}
		return nil
	}); err != nil {
		return nil, err
	}

	categories := make(category.Categories, len(categoryModels))
	for i := range categories {
		categories[i] = category.UnmarshalFromDB(categoryModels[i].ID, categoryModels[i].Name)
	}
	return categories, nil
}
