package handlers

import (
	"context"
	"fmt"
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
	"github.com/Braly-Ltd/voice-changer-api-core/ports"
	"github.com/golibs-starter/golib/log"
	"github.com/hibiken/asynq"
	"github.com/vmihailenco/msgpack/v5"
)

type VoiceChangeHandler struct {
	objectStoragePort ports.ObjectStoragePort
}

func NewVoiceChangeHandler(objectStoragePort ports.ObjectStoragePort) *VoiceChangeHandler {
	return &VoiceChangeHandler{
		objectStoragePort: objectStoragePort,
	}
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

	if err := r.objectStoragePort.DownloadFile(ctx, payload.SrcFileName); err != nil {
		return err
	}

	if err := r.objectStoragePort.UploadFilePath(ctx, payload.SrcFileName, payload.TargetFileName); err != nil {
		return err
	}

	log.Infoc(ctx, "task %s is done", payload.ID())
	return nil
}
