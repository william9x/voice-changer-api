package usecases

import (
	"context"
	"fmt"
	"github.com/Braly-Ltd/voice-changer-api-core/constants"
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
	"github.com/Braly-Ltd/voice-changer-api-core/ports"
	"github.com/Braly-Ltd/voice-changer-api-public/resources"
	"github.com/google/uuid"
)

type ChangeVoiceUseCase interface {
	CreateChangeVoiceTask(ctx context.Context, srcFile *entities.File, model string, transpose int) (*resources.Inference, error)
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

// CreateChangeVoiceTask is a use case that changes voice of an audio file.
// 1. Upload audio file to MinIO
// 2. Create a task
func (uc *ChangeVoiceUseCaseImpl) CreateChangeVoiceTask(
	ctx context.Context, srcFile *entities.File, model string, transpose int,
) (*resources.Inference, error) {
	taskId, err := uuid.NewV7()
	if err != nil {
		return nil, fmt.Errorf("generate task id error: %v", err)
	}

	taskIdStr := taskId.String()

	srcFile.Name = fmt.Sprintf("source/%s%s", taskIdStr, srcFile.Ext)
	if err := uc.objectStoragePort.UploadFile(ctx, srcFile); err != nil {
		return nil, err
	}

	targetFileName := fmt.Sprintf("target/%s%s", taskIdStr, srcFile.Ext)
	task := entities.NewVoiceChangeTask(
		taskIdStr,
		srcFile.Name,
		targetFileName,
		model,
		transpose,
		constants.TaskTypeInfer,
		constants.QueueTypeDefault,
	)

	if err := uc.taskQueuePort.Enqueue(ctx, task); err != nil {
		return nil, err
	}

	srcFileURL, err := uc.objectStoragePort.GetPreSignedObject(ctx, srcFile.Name)
	if err != nil {
		return nil, fmt.Errorf("get pre-signed src object error: %v", err)
	}

	targetFileURL, err := uc.objectStoragePort.GetPreSignedObject(ctx, targetFileName)
	if err != nil {
		return nil, fmt.Errorf("get pre-signed target object error: %v", err)
	}

	return resources.NewInferenceResource(
		taskIdStr,
		srcFileURL,
		targetFileURL,
	), nil
}
