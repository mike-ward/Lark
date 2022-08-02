package twitter

import (
	services "Lark/Services"
	"github.com/dghubble/oauth1"
	twAuth "github.com/dghubble/oauth1/twitter"
)

const (
	consumerKey    = "ZScn2AEIQrfC48Zlw"
	consumerSecret = "8gKdPBwUfZCQfUiyeFeEwVBQiV3q50wIOrIjoCxa2Q"
)

var config = oauth1.Config{
	ConsumerKey:    consumerKey,
	ConsumerSecret: consumerSecret,
	CallbackURL:    "oob",
	Endpoint:       twAuth.AuthorizeEndpoint,
}

func LogIn() (requestToken string, err error) {
	requestToken, _, err = config.RequestToken()

	if err != nil {
		return "", err
	}

	authorizationURL, err := config.AuthorizationURL(requestToken)

	if err != nil {
		return "", err
	}

	services.OpenBrowser(authorizationURL.String())
	return requestToken, err
}

func ReceivePIN(requestToken string, verifier string) (*oauth1.Token, error) {
	// Twitter ignores the oauth_signature on the access token request. The user
	// to which the request (temporary) token corresponds is already known on the
	// server. The request for a request token earlier was validated signed by
	// the consumer. Consumer applications can avoid keeping request token state
	// between authorization granting and callback handling.
	accessToken, accessSecret, err := config.AccessToken(requestToken, "secret does not matter", verifier)

	if err != nil {
		return nil, err
	}

	return oauth1.NewToken(accessToken, accessSecret), err
}
