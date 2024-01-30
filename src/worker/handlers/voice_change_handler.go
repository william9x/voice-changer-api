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
	var payload entities.VoiceChangeTask
	if err := msgpack.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("unpack task %s failed: %v", payload.ID(), err)
	}
	log.Infoc(ctx, "task %s is processing", payload.ID())

	localSourcePath := fmt.Sprintf("%s/%s", r.fileProps.BaseInputPath, payload.SrcFileName)
	if err := r.objectStoragePort.DownloadFile(ctx, payload.SrcFileName, localSourcePath); err != nil {
		return err
	}

	localTargetPath := fmt.Sprintf("%s/%s", r.fileProps.BaseOutputPath, payload.TargetFileName)

	modelPath := fmt.Sprintf("%s/%s/G.pth", r.fileProps.BaseModelPath, payload.Model)
	modelConfigPath := fmt.Sprintf("%s/%s/config.json", r.fileProps.BaseModelPath, payload.Model)
	if err := r.inferencePort.CreateInference(ctx,
		localSourcePath,
		localTargetPath,
		modelPath,
		modelConfigPath,
		payload.Transpose,
	); err != nil {
		return err
	}

	if err := r.objectStoragePort.UploadFilePath(ctx, localTargetPath, payload.TargetFileName); err != nil {
		return err
	}

	log.Infoc(ctx, "task %s is done", payload.ID())
	return nil
}
