package usecases

import (
	"context"
	"github.com/Braly-Ltd/voice-changer-api-core/ports"
	"github.com/hibiken/asynq"
)

type GetInferenceInfoUseCase interface {
	GetInferenceInfo(ctx context.Context, id string) (*asynq.TaskInfo, error)
}

type GetInferenceInfoUseCaseImpl struct {
	taskQueuePort ports.TaskQueuePort
}

func NewGetInferenceInfoUseCaseImpl(taskQueuePort ports.TaskQueuePort) *GetInferenceInfoUseCaseImpl {
	return &GetInferenceInfoUseCaseImpl{
		taskQueuePort: taskQueuePort,
	}
}

// GetInferenceInfo ...
func (uc *GetInferenceInfoUseCaseImpl) GetInferenceInfo(ctx context.Context, id string) (*asynq.TaskInfo, error) {
	return uc.taskQueuePort.GetTask(ctx, "default", id)
}
