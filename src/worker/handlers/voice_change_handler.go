package handlers

import (
	"context"
	"fmt"
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
	"github.com/golibs-starter/golib/log"
	"github.com/hibiken/asynq"
	"github.com/vmihailenco/msgpack/v5"
)

type VoiceChangeHandler struct {
}

func NewVoiceChangeHandler() *VoiceChangeHandler {
	return &VoiceChangeHandler{}
}

func (r *VoiceChangeHandler) Handle(ctx context.Context, task *asynq.Task) error {
	var payload entities.VoiceChangeTask
	if err := msgpack.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("unpack task %s failed: %v", payload.ID(), err)
	}

	log.Infoc(ctx, "task %s is done", payload.ID())
	return nil
}
