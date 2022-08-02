package main

import (
	"Lark/models"
	"Lark/resources"
	"Lark/views"
	"fyne.io/fyne/v2/app"
)

func main() {
	models.App = app.NewWithID("Lark-O-Matic")
	models.App.SetIcon(resources.AppIcon)

	models.LoadAppSettings()
	go models.TimelineFetch()

	mainWindow := views.MainWindow(models.App)
	mainWindow.SetOnClosed(func() { models.SaveAppSettings() })

	mainWindow.ShowAndRun()
}
