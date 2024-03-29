package properties

import (
	"github.com/golibs-starter/golib/config"
)

func NewFileProperties(loader config.Loader) (*FileProperties, error) {
	props := FileProperties{}
	err := loader.Bind(&props)
	return &props, err
}

type FileProperties struct {
	BaseInputPath     string
	BaseOutputPath    string
	BaseAICOutputPath string
}

func (r *FileProperties) Prefix() string {
	return "app.files"
}
