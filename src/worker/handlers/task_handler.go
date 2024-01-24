package handlers

import (
	"context"
	"github.com/hibiken/asynq"
	"go.uber.org/fx"
)

type TaskHandler interface {
	Handle(ctx context.Context, task *asynq.Task) error
}

func ProvideHandler(handlerImpl interface{}) fx.Option {
	return fx.Provide(fx.Annotate(
		handlerImpl,
		fx.As(new(TaskHandler)),
		fx.ResultTags(`group:"task_handlers"`),
	))
}
