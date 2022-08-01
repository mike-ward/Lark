package main

import (
	"Lark/models"
	"Lark/views"
	"fyne.io/fyne/v2/app"
)

func main() {
	models.App = app.NewWithID("Lark-O-Matic")
	models.LoadAppSettings()

	mainWindow := views.MainWindow(models.App)
	mainWindow.SetOnClosed(func() {
		models.SaveAppSettings()
	})

	mainWindow.ShowAndRun()
}
