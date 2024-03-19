package ports

import (
	"context"
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
)

type VoiceChangerPort interface {
	Infer(ctx context.Context, cmd entities.InferenceCommand) error
}

type AudioSeparatorPort interface {
	Infer(ctx context.Context, cmd entities.SeparateAudioCommand) (entities.SeparateAudioResponse, error)
}
