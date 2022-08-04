package views

import (
	"Lark/services"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/dghubble/go-twitter/twitter"
)

func newTweetTextContainer() *fyne.Container {
	label := widget.NewLabel("")
	label.Wrapping = fyne.TextWrapWord
	return container.NewMax(label)
}

func updateTweetText(c *fyne.Container, tweet *twitter.Tweet) {
	label := c.Objects[0].(*widget.Label)
	label.SetText(services.HtmlDecode(tweet.FullText))
}
