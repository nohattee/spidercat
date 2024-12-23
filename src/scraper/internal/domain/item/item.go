package item

import "github.com/nohattee/spidercat/src/gopkg/ulid"

type Item struct {
	id         string
	name       string
	externalID string
}

func New(externalID, name string) *Item {
	return &Item{
		id:         ulid.New(),
		name:       name,
		externalID: externalID,
	}
}

func (i *Item) ID() string {
	return i.id
}

func (i *Item) Name() string {
	return i.name
}

func (i *Item) ExternalID() string {
	return i.externalID
}
