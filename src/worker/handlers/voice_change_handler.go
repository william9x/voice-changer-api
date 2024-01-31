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

type VoiceChangeHandler struct {
	objectStoragePort ports.ObjectStoragePort
	inferencePort     ports.InferencePort
	fileProps         *properties.FileProperties
}

func NewVoiceChangeHandler(
	objectStoragePort ports.ObjectStoragePort,
	inferencePort ports.InferencePort,
	fileProps *properties.FileProperties,
) *VoiceChangeHandler {
	return &VoiceChangeHandler{
		objectStoragePort: objectStoragePort,
		inferencePort:     inferencePort,
		fileProps:         fileProps,
	}
}

func (r *VoiceChangeHandler) Type() constants.TaskType {
	return constants.TaskTypeInfer
}

// Handle
// 1. Download file from MinIO
// 2. Process file
// 3. Upload processed file to MinIO
func (r *VoiceChangeHandler) Handle(ctx context.Context, task *asynq.Task) error {
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

	modelPath := fmt.Sprintf("%s/%s/G.pth", r.fileProps.BaseModelPath, vcPayload.Model)
	modelConfigPath := fmt.Sprintf("%s/%s/config.json", r.fileProps.BaseModelPath, vcPayload.Model)
	if err := r.inferencePort.CreateInference(ctx,
		localSourcePath,
		localTargetPath,
		modelPath,
		modelConfigPath,
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
