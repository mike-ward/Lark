package views

import (
	"Lark/models"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
)

func init() {
}

func MainWindow(app fyne.App) fyne.Window {
	window := app.NewWindow("Lark")
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(300, 500))

	models.LarkSettings.IsAuthenticated.AddListener(binding.NewDataListener(func() {
		isAuthenticated, _ := models.LarkSettings.IsAuthenticated.Get()

		if isAuthenticated {
			window.SetContent(container.NewPadded(GetMainViewContainer()))
		} else {
			window.SetContent(container.NewPadded(GetPinContainer()))
		}
	}))

	return window
}
