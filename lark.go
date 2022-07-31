package main

import (
	"Lark/views"

	"fyne.io/fyne/v2/app"
)

func main() {
	application := app.NewWithID("org.mike-ward.net.lark.preferences")
	mainWindow := views.MainWindow(application)
	mainWindow.ShowAndRun()
}
