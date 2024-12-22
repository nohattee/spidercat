package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/nohattee/spidercat/src/scraper/internal/domain/author"
	"github.com/nohattee/spidercat/src/scraper/internal/domain/category"
	"github.com/nohattee/spidercat/src/scraper/internal/domain/item"
	"github.com/nohattee/spidercat/src/scraper/internal/domain/source"
	"github.com/nohattee/spidercat/src/scraper/internal/domain/tag"

	"github.com/gocolly/colly/v2"
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

func (uc *UseCase) ScrapeURLs(ctx context.Context, urls []string) error {
	mapSourceByURL := map[string]*source.Source{}
	for _, url := range urls {
		s, err := source.Parse(url)
		if err != nil {
			return err
		}
		mapSourceByURL[url] = s
	}

	for _, url := range urls {
		s := mapSourceByURL[url]
		parser := s.HTMLParser()

		c := colly.NewCollector()
		itemCollector := c.Clone()

		c.OnHTML(parser.Items(), func(e *colly.HTMLElement) {
			itemURL := e.Request.AbsoluteURL(e.Attr("href"))
			err := itemCollector.Visit(itemURL)
			if err != nil {
				log.Println(err)
			}
			panic("asd")
		})

		itemCollector.OnHTML(parser.Item(), func(e *colly.HTMLElement) {
			var err error

			tagNames := e.ChildTexts(parser.Tags())
			tags, err := uc.tagRepo.GetOrCreateByNames(ctx, tagNames)
			if err != nil {
				log.Printf("cannot get_or_create_by_names tags: %v", err)
			}

			categoryNames := e.ChildTexts(parser.Categories())
			categories, err := uc.categoryRepo.GetOrCreateByNames(ctx, categoryNames)
			if err != nil {
				log.Printf("cannot get_or_create_by_names categories: %v", err)
			}

			authorNames := e.ChildTexts(parser.Authors())
			authors, err := uc.authorRepo.GetOrCreateByNames(ctx, authorNames)
			if err != nil {
				log.Printf("cannot get_or_create_by_names authors: %v", err)
			}

			// TODO: handle images

			externalID := e.ChildText(parser.ID())
			name := e.ChildText(parser.Name())
			itemAggregate := item.NewAggregate(item.New(externalID, name), authors, categories, tags)
			err = uc.itemRepo.UpsertByExternalID(ctx, itemAggregate)
			if err != nil {
				log.Printf("cannot upsert item: %v", err)
			}
		})

		// c.OnHTML(parser.Paginator(), func(e *colly.HTMLElement) {
		// 	link := e.Attr("href")
		// 	e.Request.Visit(link)
		// })
		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL.String())
		})
		c.Visit(url)
	}

	return nil
}
