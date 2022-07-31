package models

type Settings struct {
	AccessToken       string
	AccessTokenSecret string
	ScreenName        string
}

func IsAuthenticated(settings Settings) bool {
	return settings.AccessToken != "" && settings.AccessTokenSecret != ""
}
