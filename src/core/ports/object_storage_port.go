package ports

import (
	"context"
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
)

type ObjectStoragePort interface {
	DownloadFile(ctx context.Context, name string) error
	UploadFile(ctx context.Context, object *entities.File) error
	UploadFilePath(ctx context.Context, srcName, targetName string) error
}
