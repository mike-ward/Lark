package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func GetWriteContainer() *fyne.Container {
	return container.NewMax(
		widget.NewLabelWithStyle("write tweets here", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}))
}
