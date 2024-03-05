package handlers

import (
	"context"
	"fmt"
	"github.com/Braly-Ltd/voice-changer-api-core/constants"
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
	"github.com/Braly-Ltd/voice-changer-api-core/ports"
	"github.com/Braly-Ltd/voice-changer-api-worker/properties"
	"github.com/golibs-starter/golib/log"
	"github.com/hibiken/asynq"
	"github.com/vmihailenco/msgpack/v5"
)

type RVCVoiceChangeHandler struct {
	objectStoragePort ports.ObjectStoragePort
	inferencePort     ports.InferencePort
	fileProps         *properties.FileProperties
}

func NewRVCVoiceChangeHandler(
	objectStoragePort ports.ObjectStoragePort,
	inferencePort ports.InferencePort,
	fileProps *properties.FileProperties,
) *RVCVoiceChangeHandler {
	return &RVCVoiceChangeHandler{
		objectStoragePort: objectStoragePort,
		inferencePort:     inferencePort,
		fileProps:         fileProps,
	}
}

func (r *RVCVoiceChangeHandler) Type() constants.TaskType {
	return constants.TaskTypeVoiceChangeRVC
}

// Handle
// 1. Download file from MinIO
// 2. Process file
// 3. Upload processed file to MinIO
func (r *RVCVoiceChangeHandler) Handle(ctx context.Context, task *asynq.Task) error {
	var vcPayload entities.VoiceChangePayload
	if err := msgpack.Unmarshal(task.Payload(), &vcPayload); err != nil {
		return fmt.Errorf("unpack task failed: %v", err)
	}
	log.Infoc(ctx, "task %s is processing", task.Type())

	localSourcePath := fmt.Sprintf("%s/%s", r.fileProps.BaseInputPath, vcPayload.SrcFileName)
	if err := r.objectStoragePort.DownloadFile(ctx, vcPayload.SrcFileName, localSourcePath); err != nil {
		return err
	}

	localTargetPath := fmt.Sprintf("%s/%s", r.fileProps.BaseOutputPath, vcPayload.TargetFileName)

	basePath := fmt.Sprintf("%s/%s", r.fileProps.BaseModelPath, vcPayload.Model)
	if err := r.inferencePort.CreateInference(ctx,
		localSourcePath,
		localTargetPath,
		fmt.Sprintf("%s/G.pth", basePath),
		fmt.Sprintf("%s/model.index", basePath),
		vcPayload.Transpose,
	); err != nil {
		return err
	}

	if err := r.objectStoragePort.UploadFilePath(ctx, localTargetPath, vcPayload.TargetFileName); err != nil {
		return err
	}

	log.Infoc(ctx, "task %s is done", task.Type())
	return nil
}
