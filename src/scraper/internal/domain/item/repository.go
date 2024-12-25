package item

import (
	"context"
)

type Repository interface {
	UpsertScrapedItemByExternalID(context.Context, *ScrapedItem) error
	UpsertByExternalID(context.Context, *Aggregate) error
}
