package usecase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nohattee/spidercat/src/scraper/internal/domain/author"
	"github.com/nohattee/spidercat/src/scraper/internal/domain/category"
	"github.com/nohattee/spidercat/src/scraper/internal/domain/item"
	"github.com/nohattee/spidercat/src/scraper/internal/domain/source"
	"github.com/nohattee/spidercat/src/scraper/internal/domain/tag"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
)

func (uc *UseCase) ScrapeURLs(ctx context.Context, urls []string) error {
	mapURLsBySource := map[*source.Source][]string{}
	for _, url := range urls {
		s, err := source.Parse(url)
		if err != nil {
			return err
		}
		mapURLsBySource[s] = append(mapURLsBySource[s], url)
	}

	for s, urls := range mapURLsBySource {
		parser := s.HTMLParser()
		c := colly.NewCollector(
			colly.Debugger(&debug.LogDebugger{}),
		)

		err := c.Limit(&colly.LimitRule{
			DomainGlob: "*",
			Delay:      5 * time.Second,
		})
		if err != nil {
			return fmt.Errorf("cannot set delay collector: %v", err)
		}

		itemCollector := c.Clone()

		c.OnHTML(parser.Items(), func(e *colly.HTMLElement) {
			itemURL := e.Request.AbsoluteURL(e.Attr("href"))
			err := itemCollector.Visit(itemURL)
			if err != nil {
				log.Println(err)
			}
			panic("asd")
		})

		itemCollector.OnHTML("body", func(e *colly.HTMLElement) {
			var err error

			var tags tag.Tags
			if parser.Tags() != "" {
				tagNames := e.ChildTexts(parser.Tags())
				tags, err = uc.tagRepo.GetOrCreateByNames(ctx, tagNames)
				if err != nil {
					log.Printf("cannot get_or_create_by_names tags: %v", err)
				}
			}

			var categories category.Categories
			if parser.Categories() != "" {
				categoryNames := e.ChildTexts(parser.Categories())
				categories, err = uc.categoryRepo.GetOrCreateByNames(ctx, categoryNames)
				if err != nil {
					log.Printf("cannot get_or_create_by_names categories: %v", err)
				}
			}

			var authors author.Authors
			if parser.Authors() != "" {
				authorNames := e.ChildTexts(parser.Authors())
				authors, err = uc.authorRepo.GetOrCreateByNames(ctx, authorNames)
				if err != nil {
					log.Printf("cannot get_or_create_by_names authors: %v", err)
				}
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

		for _, url := range urls {
			c.Visit(url)
		}
	}

	return nil
}
