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
	chapters     string
	imageURLs    string
	sourceID     string
	sourceURL    string
}

func NewScrapedItem(externalID, tilte, description, thumbnailURL, genres, authors, tags, chapters, imageURLs, sourceID, sourceURL string) *ScrapedItem {
	return &ScrapedItem{
		id:           ulid.New(),
		externalID:   externalID,
		tilte:        tilte,
		description:  description,
		thumbnailURL: thumbnailURL,
		genres:       genres,
		authors:      authors,
		tags:         tags,
		chapters:     chapters,
		imageURLs:    imageURLs,
		sourceID:     sourceID,
		sourceURL:    sourceURL,
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

func (i *ScrapedItem) Chapters() string {
	return i.chapters
}

func (i *ScrapedItem) ImageURLs() string {
	return i.imageURLs
}

func (i *ScrapedItem) SourceID() string {
	return i.sourceID
}

func (i *ScrapedItem) SourceURL() string {
	return i.sourceURL
}
