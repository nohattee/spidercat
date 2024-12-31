package item

import "github.com/nohattee/spidercat/src/gopkg/ulid"

type ScrapedItem struct {
	id           string
	externalID   string
	tilte        string
	description  string
	thumbnailURL string
	genres       string
	authors      string
	tags         string
	sourceID     string
	url          string
}

func NewScrapedItem(externalID, tilte, description, thumbnailURL, genres, authors, tags, sourceID, itemURL string) *ScrapedItem {
	return &ScrapedItem{
		id:           ulid.New(),
		externalID:   externalID,
		tilte:        tilte,
		description:  description,
		thumbnailURL: thumbnailURL,
		genres:       genres,
		authors:      authors,
		tags:         tags,
		sourceID:     sourceID,
		url:          itemURL,
	}
}

func (i *ScrapedItem) ID() string {
	return i.id
}

func (i *ScrapedItem) ExternalID() string {
	return i.externalID
}

func (i *ScrapedItem) Title() string {
	return i.tilte
}

func (i *ScrapedItem) Description() string {
	return i.tilte
}

func (i *ScrapedItem) ThumbnailURL() string {
	return i.tilte
}

func (i *ScrapedItem) Genres() string {
	return i.genres
}

func (i *ScrapedItem) Authors() string {
	return i.authors
}

func (i *ScrapedItem) Tags() string {
	return i.tags
}

func (i *ScrapedItem) SourceID() string {
	return i.sourceID
}

func (i *ScrapedItem) URL() string {
	return i.url
}
