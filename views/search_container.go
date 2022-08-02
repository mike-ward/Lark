package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func GetSearchContainer() *fyne.Container {
	return container.NewMax(
		widget.NewLabelWithStyle("I'm a search timeline", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))
}
