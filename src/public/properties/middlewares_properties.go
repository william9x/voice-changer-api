package properties

import (
	"github.com/golibs-starter/golib/config"
)

func NewMiddlewaresProperties(loader config.Loader) (*MiddlewaresProperties, error) {
	props := MiddlewaresProperties{}
	err := loader.Bind(&props)
	return &props, err
}

type MiddlewaresProperties struct {
	AuthenticationEnable bool `default:"true"`
}

func (t *MiddlewaresProperties) Prefix() string {
	return "app.middlewares"
}
