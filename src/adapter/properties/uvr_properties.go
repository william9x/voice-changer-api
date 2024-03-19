package properties

import (
	"fmt"
	"github.com/golibs-starter/golib/config"
)

type UVRProperties struct {
	Endpoint  string
	InferPath string

	InferURL string `default:""`
}

func NewUVRProperties(loader config.Loader) (*UVRProperties, error) {
	props := UVRProperties{}
	err := loader.Bind(&props)
	return &props, err
}

func (r *UVRProperties) Prefix() string {
	return "app.uvr"
}

func (r *UVRProperties) PostBinding() error {
	r.InferURL = fmt.Sprintf("%s%s", r.Endpoint, r.InferPath)
	return nil
}
