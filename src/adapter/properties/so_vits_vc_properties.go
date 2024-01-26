package properties

import (
	"fmt"
	"github.com/golibs-starter/golib/config"
)

type SoVitsVcProperties struct {
	Endpoint  string
	InferPath string

	InferURL string `default:""`
}

func NewSoVitsVcProperties(loader config.Loader) (*SoVitsVcProperties, error) {
	props := SoVitsVcProperties{}
	err := loader.Bind(&props)
	return &props, err
}

func (r *SoVitsVcProperties) Prefix() string {
	return "app.svc"
}

func (r *SoVitsVcProperties) PostBinding() error {
	r.InferURL = fmt.Sprintf("%s%s", r.Endpoint, r.InferPath)
	return nil
}
