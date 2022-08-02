package views

import (
	"Lark/models"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/dghubble/go-twitter/twitter"
)

func GetMainViewContainer() *fyne.Container {
	var homeTweets []twitter.Tweet
	var likeTweets []twitter.Tweet

	contentContainer := container.NewMax()

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() { updateContentContainer(contentContainer, GetTimelineContainer(homeTweets)) }),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ConfirmIcon(), func() { updateContentContainer(contentContainer, GetTimelineContainer(likeTweets)) }),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.SearchIcon(), func() { updateContentContainer(contentContainer, GetSearchContainer()) }),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.SettingsIcon(), func() { updateContentContainer(contentContainer, GetSettingsContainer()) }),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() { updateContentContainer(contentContainer, GetWriteContainer()) }),
	)

	top := container.NewVBox(
		toolbar,
		widget.NewSeparator())

	go func() {
		for {
			homeTweets = <-models.HomeTweets
			updateContentContainer(contentContainer, GetTimelineContainer(homeTweets))
			contentContainer.Refresh()
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
