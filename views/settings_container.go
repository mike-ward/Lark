package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func GetSettingsContainer() *fyne.Container {
	return container.NewMax(
		widget.NewLabelWithStyle("I'm a settings", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))
}
