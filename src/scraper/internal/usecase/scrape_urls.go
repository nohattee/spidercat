package usecase

import (
	"context"
	"fmt"
	"log"
	"net/url"
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
		chapterCollector := c.Clone()

		chapterCollector.OnHTML("body", func(e *colly.HTMLElement) {
			imageURLs := strings.Join(e.ChildAttrs(parser.ImageURLs(), "src"), ",")
			itemID := e.Request.Ctx.Get("itemID")
			chapterID := e.Request.Ctx.Get("chapterID")
			chapter := item.NewScrapedItemChapter(itemID, chapterID, e.Request.URL.String(), imageURLs)
			err = uc.itemRepo.UpsertScrapedItemChapter(ctx, chapter)
			if err != nil {
				log.Printf("cannot upsert item_chapter, err: %v, chapter: %v", err, chapter)
			}
		})

		c.OnHTML(parser.Items(), func(e *colly.HTMLElement) {
			itemURL := e.Request.AbsoluteURL(e.Attr("href"))
			err := itemCollector.Visit(itemURL)
			if err != nil {
				log.Println(err)
			}
		})

		c.OnHTML(parser.Paginators(), func(e *colly.HTMLElement) {
			url := e.Request.AbsoluteURL(e.Attr("href"))
			err := e.Request.Visit(url)
			if err != nil {
				log.Println(err)
			}
		})

		itemCollector.OnHTML("body", func(e *colly.HTMLElement) {
			var err error

			externalID := e.ChildAttr(parser.ExternalID(), "data-id")
			if externalID == "" {
				externalID = e.ChildAttr(parser.ExternalID(), "value")
			}
			title := e.ChildText(parser.Title())
			description := e.ChildText(parser.Description())
			thumbnailURL := e.ChildAttr(parser.ThumbnailURL(), "src")

			tags := strings.Join(e.ChildTexts(parser.Tags()), ",")
			genres := strings.Join(e.ChildTexts(parser.Genres()), ",")
			authors := strings.Join(e.ChildTexts(parser.Authors()), ",")

			item := item.NewScrapedItem(externalID, title, description, thumbnailURL, genres, authors, tags, s.ID(), e.Request.URL.String())
			err = uc.itemRepo.UpsertScrapedItemByExternalID(ctx, item)
			if err != nil {
				log.Printf("cannot upsert item: %v", err)
			}

			if parser.Chapters() != "" {
				chapterURLs := e.ChildAttrs(parser.Chapters(), "href")
				e.Response.Ctx.Put("itemID", item.ID())
				for i, chapterURL := range chapterURLs {
					u, err := url.Parse(chapterURL)
					if err != nil {
						log.Printf("cannot parse chapter_url: %v", err)
						continue
					}
					if u.Scheme == "" {
						chapterURL = fmt.Sprintf("https://%s%s", s.Domain(), chapterURL)
					}

					e.Response.Ctx.Put("chapterID", fmt.Sprint(i+1))
					err = chapterCollector.Request("GET", chapterURL, nil, e.Response.Ctx, nil)
					if err != nil {
						log.Println(err)
					}
				}
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
