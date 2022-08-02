package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func GetMainViewContainer() *fyne.Container {
	contentContainer := container.NewMax()
	updateContentContainer(contentContainer, GetTimelineContainer())

	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.HomeIcon(), func() { updateContentContainer(contentContainer, GetTimelineContainer()) }),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.ConfirmIcon(), func() { updateContentContainer(contentContainer, GetTimelineContainer()) }),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.SearchIcon(), func() { updateContentContainer(contentContainer, GetSearchContainer()) }),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.SettingsIcon(), func() { updateContentContainer(contentContainer, GetSettingsContainer()) }),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() { updateContentContainer(contentContainer, GetWriteContainer()) }),
	)

	return container.NewVBox(
		toolbar,
		widget.NewSeparator(),
		contentContainer,
	)
}

func updateContentContainer(contentContainer *fyne.Container, content *fyne.Container) {
	contentContainer.RemoveAll()
	contentContainer.Add(content)
}
