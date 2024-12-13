package domain

type Item struct {
	id         string
	name       string
	externalID string
	thumbnail  string
	url        string
	authors    []Author
	categories []Category
	tags       []Tag
	images     []Image
}

type Image struct {
	path string
}

type Category struct {
	id string
}

type Tag struct {
	id string
}

type Author struct {
	id string
}
