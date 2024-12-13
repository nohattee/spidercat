package usecase

import (
	"context"
	"log"

	"scraper/internal/config"
	"scraper/internal/domain/source"

	"github.com/gocolly/colly/v2"
)

type UseCase struct {
}

func New(cfg config.Config) *UseCase {
	return &UseCase{}
}

func (uc *UseCase) ScrapeURLs(ctx context.Context, urls []string) error {
	for _, url := range urls {
		s, err := source.Parse(url)
		if err != nil {
			return err
		}

		parser := s.HTMLParser()

		c := colly.NewCollector()
		itemCollector := c.Clone()

		c.OnHTML(parser.Items(), func(e *colly.HTMLElement) {
			itemURL := e.Request.AbsoluteURL(e.Attr("href"))
			itemCollector.Visit(itemURL)
		})

		c.OnHTML(parser.Item(), func(e *colly.HTMLElement) {
			var err error

			tagsName := e.ChildTexts(parser.Tags())
			tags := make(tag.Tag, len(tagsName))
			for i := range tagsName {
				tags[i] = tag.New(tagsName[i])
			}
			err = uc.repo.Tag.Upsert(ctx, tags)
			if err != nil {
				log.Printf("cannot upsert tags: %v", err)
			}

			categoriesName := e.ChildTexts(parser.Categories())
			categories := make(category.Tag, len(categoriesName))
			for i := range categoriesName {
				categories[i] = category.New(categoriesName[i])
			}
			err = uc.repo.Category.Upsert(ctx, categories)
			if err != nil {
				log.Printf("cannot upsert tags: %v", err)
			}

			authorsName := e.ChildTexts(parser.Authors())
			authors := make(author.Author, len(authorsName))
			for i := range authorsName {
				authors[i] = author.New(authorsName[i])
			}
			err = uc.repo.Author.Upsert(ctx, authors)
			if err != nil {
				log.Printf("cannot upsert authors: %v", err)
			}

			externalID := e.ChildText(parser.ID())
			name := e.ChildText(parser.Name())

			// TODO: handle images

			item := item.New(externalID, name, url, categories, authors, tags)
			err = uc.repo.Item.Upsert(ctx, item)
			if err != nil {
				log.Printf("cannot upsert item: %v", err)
			}
		})

		c.OnHTML(parser.Paginator(), func(e *colly.HTMLElement) {
			link := e.Attr("href")
			e.Request.Visit(link)
		})

		c.Visit(s.URL())
	}

	return nil
}
