package clients

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"fmt"
	"github.com/Braly-Ltd/voice-changer-api-adapter/properties"
	"google.golang.org/api/option"
)

type AuthClient struct {
	Android *auth.Client
	IOS     *auth.Client
}

func NewFirebaseAuthClient(props *properties.FirebaseProperties) (*AuthClient, error) {
	if !props.Enabled {
		return nil, nil
	}
	androidApp, err := newFirebaseApp(props.CredentialsFileAndroid)
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase app for android: %v", err)
	}

	androidAuthClient, err := androidApp.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase auth client for android: %v", err)
	}

	iosApp, err := newFirebaseApp(props.CredentialsFileIOS)
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase app for ios: %v", err)
	}

	iosAuthClient, err := iosApp.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error initializing firebase auth client for ios: %v", err)
	}

	return &AuthClient{
		Android: androidAuthClient,
		IOS:     iosAuthClient,
	}, nil
}

func newFirebaseApp(credentialsFile string) (*firebase.App, error) {
	opt := option.WithCredentialsFile(credentialsFile)
	return firebase.NewApp(context.Background(), nil, opt)
}
