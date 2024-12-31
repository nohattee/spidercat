package item

import "github.com/nohattee/spidercat/src/gopkg/ulid"

type ScrapedItemChapters []*ScrapedItemChapter

type ScrapedItemChapter struct {
	id        string
	itemID    string
	chapterID string
	url       string
	imageURLs string
}

func NewScrapedItemChapter(itemID, chapterID, chapterURL, imageURLs string) *ScrapedItemChapter {
	return &ScrapedItemChapter{
		id:        ulid.New(),
		itemID:    itemID,
		chapterID: chapterID,
		imageURLs: imageURLs,
		url:       chapterURL,
	}
}

func (i *ScrapedItemChapter) ID() string {
	return i.id
}

func (i *ScrapedItemChapter) ItemID() string {
	return i.itemID
}

func (i *ScrapedItemChapter) ChapterID() string {
	return i.chapterID
}

func (i *ScrapedItemChapter) URL() string {
	return i.url
}

func (i *ScrapedItemChapter) ImageURLs() string {
	return i.imageURLs
}
