package services

import (
	"context"
	"fmt"
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
	"github.com/Braly-Ltd/voice-changer-api-core/ports"
	"github.com/Braly-Ltd/voice-changer-api-core/utils"
	"github.com/Braly-Ltd/voice-changer-api-public/requests"
	"github.com/Braly-Ltd/voice-changer-api-public/resources"
	"github.com/hibiken/asynq"
	"time"
)

type InferenceService struct {
	objectStoragePort ports.ObjectStoragePort
	taskQueuePort     ports.TaskQueuePort
}

func NewInferenceService(
	objectStoragePort ports.ObjectStoragePort,
	taskQueuePort ports.TaskQueuePort,
) *InferenceService {
	return &InferenceService{
		objectStoragePort: objectStoragePort,
		taskQueuePort:     taskQueuePort,
	}
}

// GetInferenceInfo ...
func (r *InferenceService) GetInferenceInfo(ctx context.Context, queueId, id string) (*asynq.TaskInfo, error) {
	return r.taskQueuePort.GetTask(ctx, queueId, id)
}

func (r *InferenceService) CreateInference(ctx context.Context, req requests.CreateInferenceRequest) (resources.CreateInference, error) {
	taskId, err := utils.NewUUID()
	if err != nil {
		return resources.CreateInference{}, fmt.Errorf("generate task id error: %v", err)
	}

	srcFile := req.SrcFile
	srcFile.Name = fmt.Sprintf("source/%s%s", taskId, srcFile.Ext)
	if err := r.objectStoragePort.UploadFile(ctx, srcFile); err != nil {
		return resources.CreateInference{}, err
	}

	srcFileURL, err := r.objectStoragePort.GetPreSignedObject(ctx, srcFile.Name)
	if err != nil {
		return resources.CreateInference{}, fmt.Errorf("get pre-signed src object error: %v", err)
	}

	targetFileName := fmt.Sprintf("target/%s.mp3", taskId)
	targetFileURL, err := r.objectStoragePort.GetPreSignedObject(ctx, targetFileName)
	if err != nil {
		return resources.CreateInference{}, fmt.Errorf("get pre-signed target object error: %v", err)
	}

	packed, err := newPackedInferPayload(req, srcFile.Name, srcFileURL, targetFileName, targetFileURL)
	if err != nil {
		return resources.CreateInference{}, fmt.Errorf("pack payload error: %v", err)
	}

	queue := string(req.Queue)
	maxRetry := 0
	deadline := time.Now().Add(10 * time.Minute)
	retention := 24 * time.Hour
	taskOpts := []asynq.Option{
		asynq.TaskID(taskId),
		asynq.Queue(queue),
		asynq.MaxRetry(maxRetry),
		asynq.Deadline(deadline),
		asynq.Retention(retention),
	}

	task := asynq.NewTask(req.Type, packed, taskOpts...)
	if err := r.taskQueuePort.Enqueue(ctx, task); err != nil {
		return resources.CreateInference{}, err
	}

	return resources.CreateInference{
		ID:       utils.BuildInferenceKey(queue, taskId),
		Model:    req.Model,
		Type:     req.Type,
		Status:   asynq.TaskStatePending.String(),
		MaxRetry: maxRetry,
		Deadline: deadline.Format(time.RFC3339),

		// @Deprecated
		TaskID: utils.BuildInferenceKey(queue, taskId),
	}, nil
}

func newPackedInferPayload(
	req requests.CreateInferenceRequest,
	srFileName, srFileURL,
	targetFileName, targetFileURL string,
) ([]byte, error) {
	payload := newInferPayload(req, srFileName, srFileURL, targetFileName, targetFileURL)
	return payload.Packed()
}

func newInferPayload(
	req requests.CreateInferenceRequest,
	srFileName, srFileURL,
	targetFileName, targetFileURL string,
) *entities.VoiceChangePayload {
	return &entities.VoiceChangePayload{
		Model:          req.Model,
		Transpose:      req.Transpose,
		SrcFileName:    srFileName,
		SrcFileURL:     srFileURL,
		TargetFileName: targetFileName,
		TargetFileURL:  targetFileURL,
		EnqueuedAt:     time.Now().UnixMilli(),
	}
}
