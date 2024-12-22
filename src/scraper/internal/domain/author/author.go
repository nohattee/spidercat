package author

import "github.com/nohattee/spidercat/src/gopkg/ulid"

type Authors []*Author

type Author struct {
	id   string
	name string
}

func (a *Author) ID() string {
	return a.id
}

func (a *Author) Name() string {
	return a.name
}

func NewAuthor(name string) *Author {
	return &Author{
		id:   ulid.New(),
		name: name,
	}
}

func UnmarshalFromDB(id string, name string) *Author {
	return &Author{
		id:   id,
		name: name,
	}
}
