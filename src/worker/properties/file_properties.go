package properties

import (
	"fmt"
	"github.com/golibs-starter/golib/config"
)

func NewFileProperties(loader config.Loader) (*FileProperties, error) {
	props := FileProperties{}
	err := loader.Bind(&props)
	return &props, err
}

type FileProperties struct {
	BaseInputPath  string
	BaseOutputPath string
	BaseModelPath  string
	ModelPaths     map[string]*ModelProperties
}

type ModelProperties struct {
	Model  string
	Config string

	ModelPath  string `default:""`
	ConfigPath string `default:""`
}

func (r *FileProperties) Prefix() string {
	return "app.files"
}

func (r *FileProperties) PostBinding() error {
	for _, model := range r.ModelPaths {
		model.ModelPath = fmt.Sprintf("%s%s", r.BaseModelPath, model.Model)
		model.ConfigPath = fmt.Sprintf("%s%s", r.BaseModelPath, model.Config)
	}
	return nil
}
