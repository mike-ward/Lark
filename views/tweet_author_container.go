package views

import (
	"Lark/services"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/image/colornames"
	"html"
)

func newTweetAuthorContainer() *fyne.Container {
	screenName := canvas.NewText("", nil)
	screenName.TextStyle = fyne.TextStyle{Bold: true}

	userId := canvas.NewText("", nil)

	return container.NewHBox(
		canvas.NewText("»", colornames.Cornflowerblue),
		screenName,
		userId,
	)
}

func updateTweetAuthor(c *fyne.Container, tweet *twitter.Tweet) {
	screenName := c.Objects[1].(*canvas.Text)
	screenName.Text = html.UnescapeString(tweet.User.ScreenName)

	userId := c.Objects[2].(*canvas.Text)
	userId.Text = "@" + services.HtmlDecode(tweet.User.Name)
}
