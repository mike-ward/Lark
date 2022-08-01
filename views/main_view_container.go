package views

import (
	"Lark/views/toolbar_view"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

var _toolbarContainer *fyne.Container

func GetMainViewContainer() *fyne.Container {
	if _toolbarContainer != nil {
		return _toolbarContainer
	}

	_toolbarContainer = container.NewMax(
		toolbar_view.GetToolbarContainer())

	return _toolbarContainer
}
