package usecases

import (
	"context"
	"fmt"
	"github.com/Braly-Ltd/voice-changer-api-core/constants"
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
	"github.com/Braly-Ltd/voice-changer-api-core/ports"
	"github.com/google/uuid"
)

type ChangeVoiceUseCase interface {
	ChangeVoice(ctx context.Context, srcFile *entities.File, model string, transpose int) error
}

type ChangeVoiceUseCaseImpl struct {
	objectStoragePort ports.ObjectStoragePort
	taskQueuePort     ports.TaskQueuePort
}

func NewChangeVoiceUseCaseImpl(
	objectStoragePort ports.ObjectStoragePort,
	taskQueuePort ports.TaskQueuePort,
) *ChangeVoiceUseCaseImpl {
	return &ChangeVoiceUseCaseImpl{
		objectStoragePort: objectStoragePort,
		taskQueuePort:     taskQueuePort,
	}
}

// ChangeVoice is a use case that changes voice of an audio file.
// 1. Upload audio file to MinIO
// 2. Create a task
func (uc *ChangeVoiceUseCaseImpl) ChangeVoice(
	ctx context.Context, srcFile *entities.File, model string, transpose int,
) error {
	taskId, err := uuid.NewV7()
	if err != nil {
		return fmt.Errorf("generate task id error: %v", err)
	}

	taskIdStr := taskId.String()

	srcFile.Name = fmt.Sprintf("source/%s%s", taskIdStr, srcFile.Ext)
	if err := uc.objectStoragePort.PutObject(ctx, srcFile); err != nil {
		return err
	}

	targetFileName := fmt.Sprintf("target/%s%s", taskIdStr, srcFile.Ext)
	task := entities.NewVoiceChangeTask(
		srcFile.Name,
		targetFileName,
		model,
		transpose,
		constants.TaskTypeInfer,
		constants.QueueTypeDefault,
	)

	if err := uc.taskQueuePort.Enqueue(ctx, task); err != nil {
		return err
	}

	return nil
}
