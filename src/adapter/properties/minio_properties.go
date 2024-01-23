package properties

import (
	"github.com/golibs-starter/golib/config"
	"net/url"
)

type MinIOProperties struct {
	Endpoint          string
	PublicEndpoint    string
	publicEndpointURL url.URL
	AccessKeyID       string
	SecretAccessKey   string
	UseSSL            bool   `default:"false"`
	BucketName        string `default:"voice-changer"`
}

func NewMinIOProperties(loader config.Loader) (*MinIOProperties, error) {
	props := MinIOProperties{}
	err := loader.Bind(&props)
	return &props, err
}

func (r *MinIOProperties) Prefix() string {
	return "app.minio"
}

func (r *MinIOProperties) PostBinding() error {
	_url, err := url.Parse(r.PublicEndpoint)
	if err != nil {
		return err
	}
	r.publicEndpointURL = *_url
	return nil
}

func (r *MinIOProperties) PublicEndpointURL() url.URL {
	return r.publicEndpointURL
}
