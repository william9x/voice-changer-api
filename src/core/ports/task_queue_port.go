package ports

import (
	"context"
	"github.com/hibiken/asynq"
)

type TaskQueuePort interface {
	GetTask(ctx context.Context, queue, id string) (*asynq.TaskInfo, error)
	Enqueue(ctx context.Context, task *asynq.Task) error
}
