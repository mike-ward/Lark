package views

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func GetPinContainer() *fyne.Container {
	if _getPinContainer != nil {
		return _getPinContainer
	}

	_getPin = _askPinContainer()
	_acceptPin = _acceptPinContainer()

	_getPinContainer = container.NewPadded(
		container.NewMax(
			_getPin,
			_acceptPin,
		),
	)

	return _getPinContainer
}

/* Privates (no peeking */

var _getPin *fyne.Container
var _acceptPin *fyne.Container
var _getPinContainer *fyne.Container

func _askPinContainer() *fyne.Container {
	labelInstructions := widget.NewLabel("Clicking the button below will open the browser to get a new PIN")
	labelInstructions.Wrapping = fyne.TextWrapWord
	labelInstructions.Alignment = fyne.TextAlignCenter

	return container.NewVBox(
		labelInstructions,
		canvas.NewLine(color.Transparent),
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewButton("Get PIN", _getPinTapped),
			layout.NewSpacer(),
		),
	)
}

func _acceptPinContainer() *fyne.Container {
	labelInstructions := widget.NewLabel("Copy the PIN number from the browser into the box below and click the 'Sign in' button")
	labelInstructions.Wrapping = fyne.TextWrapWord
	labelInstructions.Alignment = fyne.TextAlignCenter

	newContainer := container.NewVBox(
		labelInstructions,
		canvas.NewLine(color.Transparent),
		container.NewHBox(
			layout.NewSpacer(),
			container.NewVBox(
				widget.NewEntry(),
				widget.NewButton("Sign in", nil),
				canvas.NewLine(color.Transparent),
				canvas.NewLine(color.Transparent),
				widget.NewButton("Start over", _startOverTapped),
			),
			layout.NewSpacer(),
		),
	)

	newContainer.Hide()
	return newContainer
}

func _getPinTapped() {
	_getPin.Hide()
	_acceptPin.Show()
}

func _startOverTapped() {
	_getPin.Show()
	_acceptPin.Hide()
}
