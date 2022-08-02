package views

import (
	"Lark/models"
	"Lark/twitter"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

func GetPinContainer() *fyne.Container {
	if _getPinContainer == nil {

		_getPin = _askPinContainer()
		_acceptPin = _acceptPinContainer()
		_inProgress = _inProgressContainer()

		_startOverTapped()

		_getPinContainer = container.NewPadded(
			container.NewMax(
				_getPin,
				_acceptPin,
				_inProgress,
			),
		)
	}

	_startOverTapped()
	return _getPinContainer
}

/* Privates (no peeking) */

var _getPin *fyne.Container
var _acceptPin *fyne.Container
var _inProgress *fyne.Container
var _getPinContainer *fyne.Container
var _pin = binding.NewString()
var _requestToken string

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

func _inProgressContainer() *fyne.Container {
	workingLabel := widget.NewLabel("Working...")
	workingLabel.Alignment = fyne.TextAlignCenter

	return container.NewVBox(
		workingLabel,
	)
}

func _acceptPinContainer() *fyne.Container {
	labelInstructions := widget.NewLabel("Copy the PIN number from the browser into the box below and click the 'Sign in' button")
	labelInstructions.Wrapping = fyne.TextWrapWord
	labelInstructions.Alignment = fyne.TextAlignCenter

	return container.NewVBox(
		labelInstructions,
		canvas.NewLine(color.Transparent),
		container.NewHBox(
			layout.NewSpacer(),
			container.NewVBox(
				widget.NewEntryWithData(_pin),
				widget.NewButton("Sign in", _signIn),
				canvas.NewLine(color.Transparent),
				canvas.NewLine(color.Transparent),
				widget.NewButton("Start over", _startOverTapped),
			),
			layout.NewSpacer(),
		),
	)
}

func _getPinTapped() {
	_getPin.Hide()
	_inProgress.Show()

	requestToken, _ := twitter.LogIn()
	_requestToken = requestToken

	_inProgress.Hide()
	_acceptPin.Show()
}

func _signIn() {
	_acceptPin.Hide()
	_inProgress.Show()

	pin, _ := _pin.Get()
	tokens, e := twitter.ReceivePIN(_requestToken, pin)

	if e != nil {
		_startOverTapped()
		return
	}

	models.UpdateAccessTokens(tokens.Token, tokens.TokenSecret)
}

func _startOverTapped() {
	_getPin.Show()
	_acceptPin.Hide()
	_inProgress.Hide()
	_requestToken = ""
	_pin.Set("")
}
