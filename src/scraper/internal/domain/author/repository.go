package author

import "context"

type Repository interface {
	GetOrCreateByNames(context.Context, []string) (Authors, error)
}
