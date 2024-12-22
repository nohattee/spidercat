package category

import "context"

type Repository interface {
	GetOrCreateByNames(context.Context, []string) (Categories, error)
}
