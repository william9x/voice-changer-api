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

type AICoverHandler struct {
	objectStoragePort ports.ObjectStoragePort
	inferencePort     ports.InferencePort
	fileProps         *properties.FileProperties
}

func NewAICoverHandler(
	objectStoragePort ports.ObjectStoragePort,
	inferencePort ports.InferencePort,
	fileProps *properties.FileProperties,
) *AICoverHandler {
	return &AICoverHandler{
		objectStoragePort: objectStoragePort,
		inferencePort:     inferencePort,
		fileProps:         fileProps,
	}
}

func (r *AICoverHandler) Type() constants.TaskType {
	return constants.TaskTypeAICover
}

// Handle
// 1. Download file from YouTube
// 2. Process file
// 3. Upload processed file to MinIO
func (r *AICoverHandler) Handle(ctx context.Context, task *asynq.Task) error {
	var vcPayload entities.VoiceChangePayload
	if err := msgpack.Unmarshal(task.Payload(), &vcPayload); err != nil {
		return fmt.Errorf("unpack task failed: %v", err)
	}
	log.Infoc(ctx, "task %s is processing", task.Type())
	log.Debugc(ctx, "task payload: %+v", vcPayload)

	localSourcePath := fmt.Sprintf("%s/%s", r.fileProps.BaseInputPath, vcPayload.SrcFileName)
	if err := r.objectStoragePort.DownloadFile(ctx, vcPayload.SrcFileName, localSourcePath); err != nil {
		return err
	}

	localTargetPath := fmt.Sprintf("%s/%s", r.fileProps.BaseOutputPath, vcPayload.TargetFileName)

	if err := r.inferencePort.CreateInference(ctx, entities.InferenceCommand{
		ModelPath: fmt.Sprintf("%s.pth", vcPayload.Model),
		IndexPath: fmt.Sprintf("%s.index", vcPayload.Model),
		InputPath: localSourcePath,
		OutPath:   localTargetPath,
		Transpose: vcPayload.Transpose,
	}); err != nil {
		return err
	}

	log.Infoc(ctx, "task %s inference completed, start uploading file at %s", task.Type(), vcPayload.TargetFileName)
	if err := r.objectStoragePort.UploadFilePath(ctx, localTargetPath, vcPayload.TargetFileName); err != nil {
		log.Errorf("upload file error: %v", err)
		return err
	}

	log.Infoc(ctx, "task %s is done", task.Type())
	return nil
}
