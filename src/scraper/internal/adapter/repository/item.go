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

func (itemModel) TableName() string {
	return "item"
}

type itemAuthorModel struct {
	ItemID    string
	AuthorID  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (itemAuthorModel) TableName() string {
	return "item_author"
}

type itemTagModel struct {
	ItemID    string
	TagID     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (itemTagModel) TableName() string {
	return "item_tag"
}

type itemCategoryModel struct {
	ItemID     string
	CategoryID string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (itemCategoryModel) TableName() string {
	return "item_category"
}

func (repo *ItemRepository) UpsertByExternalID(ctx context.Context, item *item.Aggregate) error {
	now := time.Now()

	if err := repo.db.Transaction(func(tx *gorm.DB) error {
		im := itemModel{
			ID:         item.ID(),
			Name:       item.Name(),
			ExternalID: item.ExternalID(),
			CreatedAt:  now,
			UpdatedAt:  now,
		}
		result := tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "external_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"name", "updated_at"}),
		}).Create(&im)
		if result.Error != nil {
			return result.Error
		}

		itemAuthors := make([]itemAuthorModel, len(item.Authors()))
		for i := range item.Authors() {
			author := item.Authors()[i]
			itemAuthors[i] = itemAuthorModel{
				ItemID:    im.ID,
				AuthorID:  author.ID(),
				CreatedAt: now,
				UpdatedAt: now,
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
				ItemID:    im.ID,
				TagID:     author.ID(),
				CreatedAt: now,
				UpdatedAt: now,
			}
		}

		result = tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "item_id"}, {Name: "tag_id"}},
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
				CreatedAt:  now,
				UpdatedAt:  now,
			}
		}

		result = tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "item_id"}, {Name: "category_id"}},
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
