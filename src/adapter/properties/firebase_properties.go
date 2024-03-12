package properties

import (
	"github.com/golibs-starter/golib/config"
)

type FirebaseProperties struct {
	CredentialsFileAndroid string
	CredentialsFileIOS     string
}

func NewFirebaseProperties(loader config.Loader) (*FirebaseProperties, error) {
	props := FirebaseProperties{}
	err := loader.Bind(&props)
	return &props, err
}

func (r *FirebaseProperties) Prefix() string {
	return "app.firebase"
}
