package adapter

import (
	"context"
	"firebase.google.com/go/auth"
	"github.com/Braly-Ltd/voice-changer-api-adapter/clients"
	"github.com/Braly-Ltd/voice-changer-api-core/entities"
	"github.com/golibs-starter/golib/log"
)

// FirebaseAdapter ...
type FirebaseAdapter struct {
	authClient *clients.AuthClient
}

// NewFirebaseAdapter ...
func NewFirebaseAdapter(authClient *clients.AuthClient) (*FirebaseAdapter, error) {
	return &FirebaseAdapter{
		authClient: authClient,
	}, nil
}

func (r *FirebaseAdapter) Authenticate(ctx context.Context, agent, token string) (entities.TokenData, error) {

	var tokenData *auth.Token
	var err error

	if agent == "ios" {
		tokenData, err = r.authClient.IOS.VerifyIDToken(ctx, token)
	} else {
		tokenData, err = r.authClient.Android.VerifyIDToken(ctx, token)
	}

	if err != nil {
		log.Warnf("error verifying token: %v", err)
		return entities.TokenData{}, err
	}
	return entities.TokenData{
		Issuer:   tokenData.Issuer,
		Expires:  tokenData.Expires,
		IssuedAt: tokenData.IssuedAt,
		Subject:  tokenData.Subject,
		UserID:   tokenData.UID,
		Claims:   tokenData.Claims,
	}, nil
}
