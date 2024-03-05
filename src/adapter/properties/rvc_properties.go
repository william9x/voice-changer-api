package properties

import (
	"fmt"
	"github.com/golibs-starter/golib/config"
)

type RVCProperties struct {
	Endpoint  string
	InferPath string

	InferURL string `default:""`
}

func NewRVCProperties(loader config.Loader) (*RVCProperties, error) {
	props := RVCProperties{}
	err := loader.Bind(&props)
	return &props, err
}

func (r *RVCProperties) Prefix() string {
	return "app.rvc"
}

func (r *RVCProperties) PostBinding() error {
	r.InferURL = fmt.Sprintf("%s%s", r.Endpoint, r.InferPath)
	return nil
}
