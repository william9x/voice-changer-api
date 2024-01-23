package entities

import (
	"io"
	"path"
)

type File struct {
	Name     string
	Ext      string
	Size     int64
	Content  io.Reader
	MetaData map[string]string
}

func NewFile(name string, size int64, content io.Reader) *File {
	ext := path.Ext(name)
	return &File{
		Name:     name,
		Ext:      ext,
		Size:     size,
		Content:  content,
		MetaData: make(map[string]string),
	}
}
