package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/dghubble/go-twitter/twitter"
)

func NewTimelineContainer(len int) *fyne.Container {
	vbox := container.NewVBox()
	for i := 0; i < len; i += 1 {
		vbox.Add(tweetContainer())
	}
	return container.NewMax(container.NewVScroll(vbox))
}

func UpdateTimelineContainer(c *fyne.Container, tweets *[]twitter.Tweet) {
	scroll := c.Objects[0].(*container.Scroll)
	vbox := scroll.Content.(*fyne.Container)
	len := len(*tweets)

	for i, obj := range vbox.Objects {
		tc := obj.(*fyne.Container)

		if i >= len {
			break
		}

		tweet := (*tweets)[i]
		id := tc.Objects[0].(*canvas.Text)

		if tweet.IDStr != id.Text {
			id.Text = tweet.IDStr

			screenName := tc.Objects[1].(*canvas.Text)
			screenName.Text = tweet.User.ScreenName

			text := tc.Objects[2].(*widget.Label)
			text.SetText(tweet.Text)
		}
	}
}

func tweetContainer() *fyne.Container {
	id := canvas.NewText("", nil)
	id.Hide()

	name := canvas.NewText("", nil)
	name.TextStyle = fyne.TextStyle{Bold: true}

	label := widget.NewLabel("")
	label.Wrapping = fyne.TextWrapWord

	return container.NewVBox(
		id,
		name,
		label)
}
