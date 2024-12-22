package category

import "github.com/nohattee/spidercat/src/gopkg/ulid"

type Categories []*Category

func (categories Categories) Names() []string {
	names := make([]string, len(categories))
	for i := range categories {
		names[i] = categories[i].name
	}
	return names
}

type Category struct {
	id   string
	name string
}

func (t *Category) ID() string {
	return t.id
}

func (t *Category) Name() string {
	return t.name
}

func NewCategory(name string) *Category {
	return &Category{
		id:   ulid.New(),
		name: name,
	}
}

func UnmarshalFromDB(id string, name string) *Category {
	return &Category{
		id:   id,
		name: name,
	}
}
