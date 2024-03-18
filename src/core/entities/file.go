package entities

import (
	"io"
	"mime/multipart"
	"path"
)

type File struct {
	Name     string
	Ext      string
	Size     int64
	Content  io.Reader
	MetaData map[string]string
}

func NewFile(file *multipart.FileHeader) (File, error) {
	ext := path.Ext(file.Filename)
	content, err := file.Open()
	if err != nil {
		return File{}, nil
	}
	defer content.Close()
	return File{
		Name:     file.Filename,
		Ext:      ext,
		Size:     file.Size,
		Content:  content,
		MetaData: make(map[string]string),
	}, nil
}

//func NewFileFromStream(content io.ReadCloser, size int64) (File, error) {
//	ext := path.Ext(file.Filename)
//	content, err := file.Open()
//	if err != nil {
//		return File{}, nil
//	}
//	defer content.Close()
//	return File{
//		Name:     file.Filename,
//		Ext:      ext,
//		Size:     size,
//		Content:  content,
//		MetaData: make(map[string]string),
//	}, nil
//}
