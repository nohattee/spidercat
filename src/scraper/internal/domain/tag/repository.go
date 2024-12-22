package tag

import "context"

type Repository interface {
	GetOrCreateByNames(context.Context, []string) (Tags, error)
}
