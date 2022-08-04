package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/dghubble/go-twitter/twitter"
)

func newTweetContainer() *fyne.Container {
	id := canvas.NewText("", nil)
	id.Hide()

	authorContainer := newTweetAuthorContainer()
	textContainer := newTweetTextContainer()

	return container.NewVBox(
		id,
		authorContainer,
		textContainer)
}

func updateTweetContainer(tweetContainer *fyne.Container, tweet *twitter.Tweet) {
	id := tweetContainer.Objects[0].(*canvas.Text)

	if tweet.IDStr != id.Text {
		id.Text = tweet.IDStr
		updateTweetAuthor(tweetContainer.Objects[1].(*fyne.Container), tweet)
		updateTweetText(tweetContainer.Objects[2].(*fyne.Container), tweet)
	}
}
