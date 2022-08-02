package views

import (
	"Lark/models"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/dghubble/go-twitter/twitter"
)

func GetTimelineContainer() *fyne.Container {
	vbox := container.NewVBox()

	for _, tweet := range models.HomeTweets {
		vbox.Add(tweetWidget(tweet))
	}

	return container.NewMax(container.NewVScroll(vbox))
}

func tweetWidget(tweet twitter.Tweet) *fyne.Container {
	name := canvas.NewText(tweet.User.ScreenName, nil)
	name.TextStyle = fyne.TextStyle{Bold: true}

	text := widget.NewLabel(tweet.Text)
	text.Wrapping = fyne.TextWrapWord

	return container.NewVBox(
		name,
		text)
}
