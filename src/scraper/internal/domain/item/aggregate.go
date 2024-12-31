package item

import (
	"github.com/nohattee/spidercat/src/scraper/internal/domain/author"
	"github.com/nohattee/spidercat/src/scraper/internal/domain/category"
	"github.com/nohattee/spidercat/src/scraper/internal/domain/tag"
)

type Aggregate struct {
	*Item
	authors    author.Authors
	categories category.Categories
	tags       tag.Tags
	images     Images
}

type ScrapedItemAggregate struct {
	*ScrapedItem
	chapters ScrapedItemChapters
}

func NewScrapedItemAggregate(item *ScrapedItem, chapters ScrapedItemChapters) *ScrapedItemAggregate {
	return &ScrapedItemAggregate{
		ScrapedItem: item,
		chapters:    chapters,
	}
}

func NewAggregate(item *Item, authors author.Authors, categories category.Categories, tags tag.Tags) *Aggregate {
	return &Aggregate{
		Item:       item,
		authors:    authors,
		categories: categories,
		tags:       tags,
	}
}

func (a *Aggregate) Categories() category.Categories {
	return a.categories
}

func (a *Aggregate) Authors() author.Authors {
	return a.authors
}

func (a *Aggregate) Tags() tag.Tags {
	return a.tags
}

func (a *Aggregate) Images() Images {
	return a.images
}
