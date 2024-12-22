package source

import (
	"fmt"
	"net/url"
	"strings"
)

var (
	Google = &Source{domain: GoogleDomain, htmlParser: &parser{
		id:         "",
		name:       "",
		categories: "",
		tags:       "",
		authors:    "",
		images:     "",
		thumbnail:  "",
		paginator:  "",
	}}
	TruyenQQT = &Source{domain: TruyenQQTDomain, htmlParser: &parser{
		id:         "",
		name:       "",
		categories: "",
		tags:       "",
		authors:    "",
		images:     "",
		thumbnail:  "",
		paginator:  "",
		chapters:   "",
		items:      "#main_homepage > div.list_grid_out > ul > li > div.book_avatar > a",
	}}
)

const (
	GoogleDomain    = "google.com"
	TruyenQQTDomain = "truyenqqto.com"
)

type Source struct {
	domain     string
	htmlParser *parser
}

func (s *Source) HTMLParser() *parser {
	return s.htmlParser
}

func (s *Source) Domain() string {
	return s.domain
}

type parser struct {
	id         string
	name       string
	categories string
	tags       string
	authors    string
	images     string
	thumbnail  string
	paginator  string
	item       string
	items      string
	chapters   string
}

func (p *parser) ID() string {
	return p.id
}

func (p *parser) Name() string {
	return p.name
}

func (p *parser) Paginator() string {
	return p.paginator
}

func (p *parser) Item() string {
	return p.item
}

func (p *parser) Items() string {
	return p.items
}

func (p *parser) Tags() string {
	return p.tags
}

func (p *parser) Categories() string {
	return p.categories
}

func (p *parser) Authors() string {
	return p.authors
}

func (p *parser) Chapters() string {
	return p.chapters
}

func Parse(rawURL string) (*Source, error) {
	var source *Source
	url, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	domain := strings.TrimPrefix(url.Hostname(), "www.")
	switch domain {
	case GoogleDomain:
		source = Google
	case TruyenQQTDomain:
		source = TruyenQQT
	default:
		return nil, fmt.Errorf("the '%s' domain is not supported", domain)
	}

	return source, nil
}
