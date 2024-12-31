package item

import (
	"context"
)

type Repository interface {
	UpsertScrapedItemByExternalID(context.Context, *ScrapedItem) error
	UpsertScrapedItemChapter(context.Context, *ScrapedItemChapter) error
	UpsertByExternalID(context.Context, *Aggregate) error
}
