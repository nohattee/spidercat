package item

type Images []*Image

type Image struct {
	filename string
}

func (i *Image) FileName() string {
	return i.filename
}
