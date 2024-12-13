package source

import (
	"fmt"
	"net/mail"
	"net/url"
	"strings"
)

var (
	GoogleSource = &Source{domain: GoogleDomain, htmlParser: &parser{
		id:         "",
		name:       "",
		categories: "",
		tags:       "",
		authors:    "",
		images:     "",
		thumbnail:  "",
		paginator:  "",
		items:      "",
	}}
)

const (
	GoogleDomain = "google.com"
)

type Source struct {
	domain     string
	htmlParser *parser
	jsonParser *parser
	url        string
}

func (s *Source) HTMLParser() *parser {
	return s.htmlParser
}

func (s *Source) Domain() string {
	return s.domain
}

func (s *Source) URL() string {
	return s.url
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

func Parse(rawURL string) (*Source, error) {
	var source *Source
	url, err := url.Parse(rawURL)
	if err != nil {
		return nil, err
	}
	domain := strings.TrimPrefix(url.Hostname(), "www.")
	switch domain {
	case GoogleDomain:
		source = GoogleSource
	default:
		return nil, fmt.Errorf("the %s domain is not supported", domain)
	}
	mail.ParseAddress()

	return source, nil
}
