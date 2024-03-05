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

func (r *MinIOAdapter) DownloadFile(ctx context.Context, name, destPath string) error {
	if err := r.client.FGetObject(ctx, r.props.BucketName, name, destPath, minio.GetObjectOptions{}); err != nil {
		return fmt.Errorf("download object error: %v", err)
	}
	return nil
}

// UploadFile ...
func (r *MinIOAdapter) UploadFile(ctx context.Context, object entities.File) error {
	info, err := r.client.PutObject(ctx, r.props.BucketName, object.Name, object.Content, object.Size, minio.PutObjectOptions{
		ContentType:  "application/octet-stream",
		UserMetadata: object.MetaData,
	})
	if err != nil {
		return fmt.Errorf("upload file error: %v", err)
	}

	log.Debugc(ctx, "uploaded file: bucket %s name %s", info.Bucket, info.Key)
	return nil
}

func (r *MinIOAdapter) UploadFilePath(ctx context.Context, targetFile, targetName string) error {
	info, err := r.client.FPutObject(ctx, r.props.BucketName, targetName, targetFile, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		return fmt.Errorf("upload file path error: %v", err)
	}

	log.Debugc(ctx, "uploaded file path: bucket %s name %s", info.Bucket, info.Key)
	return nil
}

func (r *MinIOAdapter) IsObjectExist(ctx context.Context, objectName string) bool {
	_, err := r.client.StatObject(ctx, r.props.BucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		log.Debugc(ctx, "bucket %s object %s get stats error: %v", r.props.BucketName, objectName, err)
		return false
	}
	return true
}

func (r *MinIOAdapter) GetPreSignedObject(ctx context.Context, objectName string) (string, error) {
	//url, err := r.client.PresignedGetObject(ctx, r.props.BucketName, objectName, 3600*time.Second, nil)
	//if err != nil {
	//	return "", fmt.Errorf("get presigned object error: %v", err)
	//}
	//return url.Path, nil
	endpoint := r.props.PublicEndpointURL()
	return endpoint.JoinPath(r.props.BucketName, objectName).String(), nil
}
