package ports

import (
	"context"
	"github.com/Braly-Ltd/voice-changer-api-core/constants"
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
)

type TaskQueuePort interface {
	Enqueue(ctx context.Context, taskType constants.TaskType, task entities.Task) error
}
