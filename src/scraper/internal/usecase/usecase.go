package usecase

import (
	"context"

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
			id := e.ChildText(parser.ID())
			name := e.ChildText(parser.Name())
			tags := e.ChildTexts(parser.Tags())
			categories := e.ChildTexts(parser.Categories())
			authors := e.ChildTexts(parser.Authors())

		})

		c.OnHTML(parser.Paginator(), func(e *colly.HTMLElement) {
			link := e.Attr("href")
			e.Request.Visit(link)
		})

		c.Visit(s.URL())
	}

	return nil
}
