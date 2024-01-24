package ports

import (
	"context"
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
)

type TaskQueuePort interface {
	Enqueue(ctx context.Context, task entities.Task) error
}
