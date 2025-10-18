package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"github.com/sharkmu/jhournal/tabs"
)

func main() {
	a := app.NewWithID("com.sharkmu.jhournal")
	w := a.NewWindow("Jhournal")
	w.Resize(fyne.NewSize(800, 600))
	w.CenterOnScreen()

	var viewTab *container.TabItem

	refreshView := func() {
		viewTab.Content = tabs.ViewEntries()
	}

	newTab := container.NewTabItem("New Entry", tabs.NewEntry(refreshView))

	viewTab = container.NewTabItem("View Entries", tabs.ViewEntries())

	settingsTab := container.NewTabItem("Settings", tabs.OpenSettings(w))

	tabsContainer := container.NewAppTabs(newTab, viewTab, settingsTab)
	tabsContainer.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabsContainer)
	w.ShowAndRun()
}
