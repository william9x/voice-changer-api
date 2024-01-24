package adapter

import (
	"context"
	"fmt"
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
	"github.com/golibs-starter/golib/log"
	"github.com/hibiken/asynq"
)

// AsynqAdapter ...
type AsynqAdapter struct {
	client *asynq.Client
}

// NewAsynqAdapter ...
func NewAsynqAdapter(client *asynq.Client) *AsynqAdapter {
	return &AsynqAdapter{client: client}
}

// Enqueue ...
func (c *AsynqAdapter) Enqueue(ctx context.Context, task entities.Task) error {
	packed, err := task.Pack()
	if err != nil {
		return fmt.Errorf("pack payload error: %v", err)
	}

	taskOpts := []asynq.Option{
		asynq.TaskID(task.ID()),
		asynq.Queue(string(task.Queue())),
	}
	info, err := c.client.EnqueueContext(ctx, asynq.NewTask(string(task.Type()), packed), taskOpts...)
	if err != nil {
		return fmt.Errorf("enqueue task error: %v", err)
	}

	log.Infof("enqueue task: id %s type %s queue %s", info.ID, info.Type, info.Queue)
	return err
}
