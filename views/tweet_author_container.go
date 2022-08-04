package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/dghubble/go-twitter/twitter"
)

func newTweetAuthorContainer() *fyne.Container {
	screenName := canvas.NewText("", nil)
	screenName.TextStyle = fyne.TextStyle{Bold: true}

	userId := canvas.NewText("", nil)

	return container.NewHBox(
		screenName,
		userId,
	)
}

func updateTweetAuthor(c *fyne.Container, tweet *twitter.Tweet) {
	screenName := c.Objects[0].(*canvas.Text)
	screenName.Text = tweet.User.ScreenName

	userId := c.Objects[1].(*canvas.Text)
	userId.Text = "@" + tweet.User.Name
}
