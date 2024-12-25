package usecase

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/nohattee/spidercat/src/scraper/internal/domain/item"
	"github.com/nohattee/spidercat/src/scraper/internal/domain/source"

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
		})

		itemCollector.OnHTML("body", func(e *colly.HTMLElement) {
			var err error

			externalID := e.ChildText(parser.ExternalID())
			title := e.ChildText(parser.Title())
			description := e.ChildText(parser.Description())
			thumbnailURL := e.ChildText(parser.ThumbnailURL())

			tags := strings.Join(e.ChildTexts(parser.Tags()), ",")
			genres := strings.Join(e.ChildTexts(parser.Genres()), ",")
			authors := strings.Join(e.ChildTexts(parser.Authors()), ",")
			chapters := strings.Join(e.ChildTexts(parser.Chapters()), ",")
			imageURLs := strings.Join(e.ChildTexts(parser.ImageURLs()), ",")

			// TODO: handle images

			item := item.NewScrapedItem(externalID, title, description, thumbnailURL, genres, authors, tags, chapters, imageURLs, s.ID(), e.Request.URL.Path)
			err = uc.itemRepo.UpsertScrapedItemByExternalID(ctx, item)
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
