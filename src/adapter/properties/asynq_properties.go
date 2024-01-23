package properties

import (
	"github.com/golibs-starter/golib/config"
)

type AsynqProperties struct {
	Addr     string `default:"localhost:6379"`
	Username string `default:""`
	Password string `default:"secret"`
	DB       int    `default:"1"`
	PoolSize int    `default:"0"`
}

func NewAsynqProperties(loader config.Loader) (*AsynqProperties, error) {
	props := AsynqProperties{}
	err := loader.Bind(&props)
	return &props, err
}

func (r *AsynqProperties) Prefix() string {
	return "app.asynq"
}
