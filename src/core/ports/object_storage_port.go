package ports

import (
	"context"
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
)

type ObjectStoragePort interface {
	DownloadFile(ctx context.Context, name, destPath string) error
	UploadFile(ctx context.Context, object entities.File) error
	UploadFilePath(ctx context.Context, srcName, targetName string) error
	IsObjectExist(ctx context.Context, objectName string) bool
	GetPreSignedObject(ctx context.Context, objectName string) (string, error)
}
