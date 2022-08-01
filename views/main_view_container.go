package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var _appTabContainer *fyne.Container

func GetMainViewContainer() *fyne.Container {
	if _appTabContainer != nil {
		return _appTabContainer
	}

	_appTabContainer = container.NewMax(MyTabContainer())
	return _appTabContainer
}

func MyTabContainer() *fyne.Container {
	radioGroup := widget.NewRadioGroup([]string{"A", "B", "C"}, nil)
	radioGroup.Horizontal = true

	return container.NewVBox(
		radioGroup,
		container.NewMax(),
	)
}
