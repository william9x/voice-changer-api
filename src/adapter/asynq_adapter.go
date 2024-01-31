package adapter

import (
	"context"
	"fmt"
	"github.com/golibs-starter/golib/log"
	"github.com/hibiken/asynq"
)

// AsynqAdapter ...
type AsynqAdapter struct {
	client    *asynq.Client
	inspector *asynq.Inspector
}

// NewAsynqAdapter ...
func NewAsynqAdapter(client *asynq.Client, inspector *asynq.Inspector) *AsynqAdapter {
	return &AsynqAdapter{client: client, inspector: inspector}
}

//func (c *AsynqAdapter) GetTask(ctx context.Context, id string) error {
//	task, err := c.inspector.GetTaskInfo("default", id)
//	if err != nil {
//		return err
//	}
//	entities.Task()
//}

// Enqueue ...
func (c *AsynqAdapter) Enqueue(ctx context.Context, task *asynq.Task) error {
	info, err := c.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("enqueue task error: %v", err)
	}

	log.Infoc(ctx, "enqueue task: id %s type %s queue %s", info.ID, info.Type, info.Queue)
	return err
}
