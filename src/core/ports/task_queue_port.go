package ports

import (
	"context"
	"github.com/hibiken/asynq"
)

type TaskQueuePort interface {
	Enqueue(ctx context.Context, task *asynq.Task) error
}
