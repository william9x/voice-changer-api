package usecases

import (
	"context"
	"fmt"
	"github.com/Braly-Ltd/voice-changer-api-core/constants"
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
	"github.com/Braly-Ltd/voice-changer-api-core/ports"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"time"
)

type ChangeVoiceUseCase interface {
	CreateChangeVoiceTask(ctx context.Context, srcFile *entities.File, model string, transpose int) (string, error)
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
	ctx context.Context,
	srcFile *entities.File,
	model string,
	transpose int,
) (string, error) {
	taskId, err := uuid.NewV7()
	if err != nil {
		return "", fmt.Errorf("generate task id error: %v", err)
	}

	taskIdStr := taskId.String()

	srcFile.Name = fmt.Sprintf("source/%s%s", taskIdStr, srcFile.Ext)
	if err := uc.objectStoragePort.UploadFile(ctx, srcFile); err != nil {
		return "", err
	}

	targetFileName := fmt.Sprintf("target/%s%s", taskIdStr, srcFile.Ext)

	srcFileURL, err := uc.objectStoragePort.GetPreSignedObject(ctx, srcFile.Name)
	if err != nil {
		return "", fmt.Errorf("get pre-signed src object error: %v", err)
	}

	targetFileURL, err := uc.objectStoragePort.GetPreSignedObject(ctx, targetFileName)
	if err != nil {
		return "", fmt.Errorf("get pre-signed target object error: %v", err)
	}

	payload := entities.NewVoiceChangePayload(srcFile.Name, srcFileURL, targetFileName, targetFileURL, model, transpose)
	packed, err := payload.Packed()
	if err != nil {
		return "", fmt.Errorf("pack payload error: %v", err)
	}

	taskOpts := []asynq.Option{
		asynq.TaskID(taskIdStr),
		asynq.Queue(string(constants.QueueTypeDefault)),
		asynq.MaxRetry(0),
		asynq.Deadline(time.Now().Add(10 * time.Minute)),
		asynq.Retention(24 * time.Hour),
	}
	task := asynq.NewTask(string(constants.TaskTypeInfer), packed, taskOpts...)
	if err := uc.taskQueuePort.Enqueue(ctx, task); err != nil {
		return "", err
	}

	return taskIdStr, nil
}
