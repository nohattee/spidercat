package source

import (
	"fmt"
	"net/url"
	"strings"
)

var (
	Google = &Source{domain: GoogleDomain, htmlParser: &parser{
		externalID:   "",
		title:        "",
		genres:       "",
		tags:         "",
		authors:      "",
		imageURLs:    "",
		thumbnailURL: "",
		paginators:   "",
	}}
	TruyenQQT = &Source{id: "truyenqqto", domain: TruyenQQTDomain, htmlParser: &parser{
		externalID:   "#book_id",
		title:        "body > div.content > div.div_middle > div.main_content > div.book_detail > div.book_info > div.book_other > h1",
		description:  "body > div.content > div.div_middle > div.main_content > div.book_detail > div.story-detail-info.detail-content",
		genres:       "body > div.content > div.div_middle > div.main_content > div.book_detail > div.book_info > div.book_other > ul.list01 > li > a",
		tags:         "",
		authors:      "body > div.content > div.div_middle > div.main_content > div.book_detail > div.book_info > div.book_other > div.txt > ul > li.author.row > p.col-xs-9 > a",
		imageURLs:    "",
		thumbnailURL: "body > div.content > div.div_middle > div.main_content > div.book_detail > div.book_info > div.book_avatar > img",
		chapters:     "body > div.content > div.div_middle > div.main_content > div.book_detail > div.list_chapter > div > div > div.col-md-10.col-sm-10.col-xs-8.name-chap > a",
		items:        "#main_homepage > div.list_grid_out > ul > li > div.book_avatar > a",
		paginators:   "#main_homepage > div.page_redirect > a",
	}}
)

const (
	GoogleDomain    = "google.com"
	TruyenQQTDomain = "truyenqqto.com"
)

type Source struct {
	id         string
	domain     string
	htmlParser *parser
}

func (s *Source) ID() string {
	return s.id
}

func (s *Source) HTMLParser() *parser {
	return s.htmlParser
}

func (s *Source) Domain() string {
	return s.domain
}

type parser struct {
	externalID   string
	title        string
	description  string
	thumbnailURL string
	genres       string
	tags         string
	authors      string
	imageURLs    string
	chapters     string

	paginators string
	items      string
}

func (p *parser) ExternalID() string {
	return p.externalID
}

func (p *parser) Title() string {
	return p.title
}

func (p *parser) Description() string {
	return p.description
}

func (p *parser) ThumbnailURL() string {
	return p.thumbnailURL
}

func (p *parser) ImageURLs() string {
	return p.imageURLs
}

func (p *parser) Paginators() string {
	return p.paginators
}

func (p *parser) Items() string {
	return p.items
}

func (p *parser) Tags() string {
	return p.tags
}

func (p *parser) Genres() string {
	return p.genres
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
