package properties

import (
	"github.com/golibs-starter/golib/config"
)

func NewWorkerProperties(loader config.Loader) (*WorkerProperties, error) {
	props := WorkerProperties{}
	err := loader.Bind(&props)
	return &props, err
}

type WorkerProperties struct {
	Concurrency    int
	StrictPriority bool `default:"true"`
}

func (r *WorkerProperties) Prefix() string {
	return "app.asynq.worker"
}
