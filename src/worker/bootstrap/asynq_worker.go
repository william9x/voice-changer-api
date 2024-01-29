package bootstrap

import (
	"context"
	"github.com/Braly-Ltd/voice-changer-api-worker/workers"
	"go.uber.org/fx"
)

func ProvideAsynqWorker() fx.Option {
	return fx.Provide(
		fx.Annotate(
			workers.NewAsynqWorker,
			fx.ParamTags(`group:"task_handlers"`),
		),
	)
}

func AsynqWorkerOpt() fx.Option {
	return fx.Invoke(func(worker *workers.AsynqWorker) {
		go worker.Start()
	})
}

func OnStopAsynqWorker() fx.Option {
	return fx.Invoke(func(lc fx.Lifecycle, worker *workers.AsynqWorker) {
		lc.Append(fx.Hook{
			OnStop: func(ctx context.Context) error {
				worker.Stop()
				return nil
			},
		})
	})
}
