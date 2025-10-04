package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"github.com/sharkmu/jhournal/tabs"
)

func main() {
	a := app.New()
	w := a.NewWindow("Jhournal")
	w.Resize(fyne.NewSize(800, 600))
	w.CenterOnScreen()

	tabs := container.NewAppTabs(
		container.NewTabItem("New entry", tabs.NewEntry()),
		container.NewTabItem("View entries", tabs.ViewEntries()),
		container.NewTabItem("Settings", tabs.OpenSettings()),
	)

	tabs.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabs)
	w.ShowAndRun()
}
