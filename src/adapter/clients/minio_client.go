package clients

import (
	"github.com/Braly-Ltd/voice-changer-api-adapter/properties"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinIOClient(props *properties.MinIOProperties) (*minio.Client, error) {
	return minio.New(props.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(props.AccessKeyID, props.SecretAccessKey, ""),
		Secure: props.UseSSL,
	})
}
