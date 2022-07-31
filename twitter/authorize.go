package twitter

import (
	"github.com/dghubble/oauth1"
	twAuth "github.com/dghubble/oauth1/twitter"
	"os/exec"
	"runtime"
)

var config = oauth1.Config{
	ConsumerKey:    "ZScn2AEIQrfC48Zlw",
	ConsumerSecret: "8gKdPBwUfZCQfUiyeFeEwVBQiV3q50wIOrIjoCxa2Q",
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

	openBrowser(authorizationURL.String())
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

// openBrowser tries to open the URL in a browser,
// and returns whether it succeed in doing so.
func openBrowser(url string) bool {
	var args []string

	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}

	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}
