package views

import (
	"fyne.io/fyne/v2"
)

func MainWindow(app fyne.App) fyne.Window {
	window := app.NewWindow("Lark")

	window.CenterOnScreen()
	window.Resize(fyne.NewSize(300, 500))

	getPin := GetPinContainer()
	window.SetContent(getPin)
	return window
}
