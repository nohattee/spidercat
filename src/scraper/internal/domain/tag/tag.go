package tag

import "github.com/nohattee/spidercat/src/gopkg/ulid"

type Tags []*Tag

type Tag struct {
	id   string
	name string
}

func (t *Tag) ID() string {
	return t.id
}

func (t *Tag) Name() string {
	return t.name
}

func NewTag(name string) *Tag {
	return &Tag{
		id:   ulid.New(),
		name: name,
	}
}

func UnmarshalFromDB(id string, name string) *Tag {
	return &Tag{
		id:   id,
		name: name,
	}
}
