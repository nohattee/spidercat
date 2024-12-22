package item

import "context"

type Repository interface {
	UpsertByExternalID(context.Context, *Aggregate) error
}
