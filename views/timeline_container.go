package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/dghubble/go-twitter/twitter"
)

func newTimelineContainer(len int) *fyne.Container {
	vbox := container.NewVBox()
	for i := 0; i < len; i += 1 {
		vbox.Add(newTweetContainer())
	}
	return container.NewMax(container.NewVScroll(vbox))
}

func updateTimeline(c *fyne.Container, tweets *[]twitter.Tweet) {
	scroll := c.Objects[0].(*container.Scroll)
	vbox := scroll.Content.(*fyne.Container)
	len := len(*tweets)

	for i, obj := range vbox.Objects {
		if i >= len {
			break
		}

		tweetContainer := obj.(*fyne.Container)
		tweet := &(*tweets)[i]
		updateTweetContainer(tweetContainer, tweet)
	}
}
