package ports

import (
	"context"
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
)

type ObjectStoragePort interface {
	PutObject(ctx context.Context, object *entities.File) error
}
