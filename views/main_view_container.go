package views

import (
	"Lark/models"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/dghubble/go-twitter/twitter"
)

const (
	HOME     int = 0
	LIKE     int = 1
	SEARCH   int = 2
	SETTINGS int = 3
	WRITE    int = 4
)

func GetMainViewContainer() *fyne.Container {
	var homeTweets *[]twitter.Tweet
	//var likeTweets *[]twitter.Tweet

	selected := HOME
	homeTimeline := NewTimelineContainer(75)
	likeTimeline := NewTimelineContainer(75)

	contentContainer := container.NewMax()

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() {
			selected = HOME
			updateContentContainer(contentContainer, homeTimeline)
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ConfirmIcon(), func() {
			selected = LIKE
			updateContentContainer(contentContainer, likeTimeline)
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.SearchIcon(), func() {
			selected = SEARCH
			updateContentContainer(contentContainer, GetSearchContainer())
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {
			selected = SETTINGS
			updateContentContainer(contentContainer, GetSettingsContainer())
		}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			selected = WRITE
			updateContentContainer(contentContainer, GetWriteContainer())
		}),
	)

	top := container.NewVBox(
		toolbar,
		widget.NewSeparator())

	go func() {
		for {
			homeTweets = <-models.HomeTweets
			if selected == HOME {
				UpdateTimelineContainer(homeTimeline, homeTweets)
				updateContentContainer(contentContainer, homeTimeline)
				contentContainer.Refresh()
			}
		}
	}()

	return container.NewBorder(
		top,
		nil,
		nil,
		nil,
		contentContainer,
	)
}

func updateContentContainer(contentContainer *fyne.Container, content *fyne.Container) {
	contentContainer.RemoveAll()
	contentContainer.Add(content)
}
