package adapter

import (
	"context"
	"fmt"
	"github.com/Braly-Ltd/voice-changer-api-core/constants"
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
func (c *AsynqAdapter) Enqueue(ctx context.Context, taskType constants.TaskType, task entities.Task) error {
	packed, err := task.Pack()
	if err != nil {
		return fmt.Errorf("pack payload error: %v", err)
	}
	info, err := c.client.EnqueueContext(ctx, asynq.NewTask(string(taskType), packed), asynq.Queue(task.Queue()))
	log.Infof("Enqueue task: %v", info)
	return err
}
