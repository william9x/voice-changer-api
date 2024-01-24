package adapter

import (
	"context"
	"fmt"
	"github.com/Braly-Ltd/voice-changer-api-adapter/properties"
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
	"github.com/golibs-starter/golib/log"
	"github.com/minio/minio-go/v7"
)

// MinIOAdapter ...
type MinIOAdapter struct {
	client *minio.Client
	props  *properties.MinIOProperties
}

// NewMinIOAdapter ...
func NewMinIOAdapter(client *minio.Client, clientProps *properties.MinIOProperties) *MinIOAdapter {
	return &MinIOAdapter{client: client, props: clientProps}
}

func (c *MinIOAdapter) DownloadFile(ctx context.Context, name string) error {
	path := "/tmp/" + name
	if err := c.client.FGetObject(ctx, c.props.BucketName, name, path, minio.GetObjectOptions{}); err != nil {
		return fmt.Errorf("download object error: %v", err)
	}
	return nil
}

// UploadFile ...
func (c *MinIOAdapter) UploadFile(ctx context.Context, object *entities.File) error {
	info, err := c.client.PutObject(ctx, c.props.BucketName, object.Name, object.Content, object.Size, minio.PutObjectOptions{
		ContentType:  "application/octet-stream",
		UserMetadata: object.MetaData,
	})
	if err != nil {
		return fmt.Errorf("upload file error: %v", err)
	}

	log.Debugc(ctx, "uploaded file: bucket %s name %s", info.Bucket, info.Key)
	return nil
}

func (c *MinIOAdapter) UploadFilePath(ctx context.Context, srcName, targetName string) error {
	sourceFile := "/tmp/" + srcName
	info, err := c.client.FPutObject(ctx, c.props.BucketName, targetName, sourceFile, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		return fmt.Errorf("upload file path error: %v", err)
	}

	log.Debugc(ctx, "uploaded file path: bucket %s name %s", info.Bucket, info.Key)
	return nil
}
