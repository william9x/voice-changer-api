package ports

import (
	"context"
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
)

type InferencePort interface {
	CreateInference(ctx context.Context, cmd entities.InferenceCommand) error
	SeperateAudio(ctx context.Context, cmd entities.SeparateAudioCommand) (entities.SeparateAudioResponse, error)
}
