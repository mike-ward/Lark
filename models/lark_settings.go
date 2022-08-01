package models

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
)

var App fyne.App

var LarkSettings struct {
	AccessToken       string
	AccessTokenSecret string
	ScreenName        string
	IsAuthenticated   binding.Bool `json:"-"`
}

func init() {
	LarkSettings.AccessToken = ""
	LarkSettings.AccessTokenSecret = ""
	LarkSettings.ScreenName = ""
	LarkSettings.IsAuthenticated = binding.NewBool()
}

func LoadAppSettings() {
	res := App.Preferences().String("LarkSettings")

	if res != "" {
		err := json.Unmarshal([]byte(res), &LarkSettings)

		if err == nil {
			UpdateAccessTokens(LarkSettings.AccessToken, LarkSettings.AccessTokenSecret)
		}
	}
}

func SaveAppSettings() {
	res, e := json.Marshal(LarkSettings)

	if e == nil {
		resString := string(res)
		App.Preferences().SetString("LarkSettings", resString)
	}
}

func UpdateAccessTokens(accessToken string, accessTokenSecret string) {
	LarkSettings.AccessToken = accessToken
	LarkSettings.AccessTokenSecret = accessTokenSecret
	LarkSettings.IsAuthenticated.Set(accessToken != "" && accessTokenSecret != "")
}
