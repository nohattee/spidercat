package repository

import (
	"context"
	"time"

	"github.com/nohattee/spidercat/src/scraper/internal/domain/item"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{
		db: db,
	}
}

type itemModel struct {
	ID         string
	Name       string
	ExternalID string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type itemAuthorModel struct {
	ItemID   string
	AuthorID string
}

type itemTagModel struct {
	ItemID string
	TagID  string
}

type itemCategoryModel struct {
	ItemID     string
	CategoryID string
}

func (repo *ItemRepository) UpsertByExternalID(ctx context.Context, item *item.Aggregate) error {
	now := time.Now()

	if err := repo.db.Transaction(func(tx *gorm.DB) error {
		im := itemModel{
			ID:         item.ID(),
			Name:       item.Name(),
			ExternalID: item.ExternalID(),
			UpdatedAt:  now,
		}
		result := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "external_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"name"}),
		}).Create(&im)
		if result.Error != nil {
			return result.Error
		}

		itemAuthors := make([]itemAuthorModel, len(item.Authors()))
		for i := range item.Authors() {
			author := item.Authors()[i]
			itemAuthors[i] = itemAuthorModel{
				ItemID:   im.ID,
				AuthorID: author.ID(),
			}
		}

		result = tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "item_id"}, {Name: "author_id"}},
			DoNothing: true,
		}).CreateInBatches(itemAuthors, 100)
		if result.Error != nil {
			return result.Error
		}

		itemTags := make([]itemTagModel, len(item.Tags()))
		for i := range item.Tags() {
			author := item.Tags()[i]
			itemTags[i] = itemTagModel{
				ItemID: im.ID,
				TagID:  author.ID(),
			}
		}

		result = tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "item_id"}, {Name: "author_id"}},
			DoNothing: true,
		}).CreateInBatches(itemTags, 100)
		if result.Error != nil {
			return result.Error
		}

		itemCategories := make([]itemCategoryModel, len(item.Categories()))
		for i := range item.Categories() {
			author := item.Categories()[i]
			itemCategories[i] = itemCategoryModel{
				ItemID:     im.ID,
				CategoryID: author.ID(),
			}
		}

		result = tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "item_id"}, {Name: "author_id"}},
			DoNothing: true,
		}).CreateInBatches(itemCategories, 100)
		if result.Error != nil {
			return result.Error
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
