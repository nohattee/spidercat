package usecase

import (
	"github.com/nohattee/spidercat/src/scraper/internal/domain/author"
	"github.com/nohattee/spidercat/src/scraper/internal/domain/category"
	"github.com/nohattee/spidercat/src/scraper/internal/domain/item"
	"github.com/nohattee/spidercat/src/scraper/internal/domain/tag"
)

type UseCase struct {
	tagRepo      tag.Repository
	categoryRepo category.Repository
	authorRepo   author.Repository
	itemRepo     item.Repository
}

func New(itemRepo item.Repository, categoryRepo category.Repository, authorRepo author.Repository, tagRepo tag.Repository) *UseCase {
	return &UseCase{
		itemRepo:     itemRepo,
		categoryRepo: categoryRepo,
		authorRepo:   authorRepo,
		tagRepo:      tagRepo,
	}
}
