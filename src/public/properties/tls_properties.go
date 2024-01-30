package properties

import (
	"github.com/golibs-starter/golib/config"
)

func NewTLSProperties(loader config.Loader) (*TLSProperties, error) {
	props := TLSProperties{}
	err := loader.Bind(&props)
	return &props, err
}

type TLSProperties struct {
	Enabled  bool `default:"false"`
	CertFile string
	KeyFile  string
}

func (t TLSProperties) Prefix() string {
	return "app.tls"
}
