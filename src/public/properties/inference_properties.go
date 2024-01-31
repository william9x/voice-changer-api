package properties

import (
	"github.com/golibs-starter/golib/config"
)

func NewInferenceProperties(loader config.Loader) (*InferenceProperties, error) {
	props := InferenceProperties{}
	err := loader.Bind(&props)
	return &props, err
}

type InferenceProperties struct {
	SupportedFiles []string
}

func (t *InferenceProperties) Prefix() string {
	return "app.inference"
}
