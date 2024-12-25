package item

import "github.com/nohattee/spidercat/src/gopkg/ulid"

type Item struct {
	id          string
	externalID  string
	title       string
	description string
}

func New(externalID, title, description string) *Item {
	return &Item{
		id:          ulid.New(),
		externalID:  externalID,
		title:       title,
		description: description,
	}
}

func (i *Item) ID() string {
	return i.id
}

func (i *Item) ExternalID() string {
	return i.externalID
}

func (i *Item) Title() string {
	return i.title
}

func (i *Item) Description() string {
	return i.description
}
