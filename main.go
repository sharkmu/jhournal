package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"github.com/sharkmu/jhournal/tabs"
	"github.com/sharkmu/jhournal/utils"
)

func main() {
	a := app.NewWithID("com.sharkmu.jhournal")
	w := a.NewWindow("Jhournal")

	wSize := utils.GetSize()
	w.Resize(wSize)

	w.CenterOnScreen()

	w.Show()

	createTabs(w)

	a.Run()
}

func createTabs(w fyne.Window) {
	var newTab, viewTab, settingsTab *container.TabItem
	var refreshTab func(tabName string)

	refreshTab = func(tabName string) {
		switch tabName {
		case "new":
			newTab.Content = tabs.NewEntry(refreshTab)
		case "view":
			viewTab.Content = tabs.ViewEntries()
		case "settings":
			settingsTab.Content = tabs.OpenSettings(w, refreshTab)
		default:
			utils.DisplayError(fmt.Errorf("no such tab to refresh: %v", tabName))
		}
	}

	newTab = container.NewTabItem("New Entry", tabs.NewEntry(refreshTab))
	viewTab = container.NewTabItem("View Entries", tabs.ViewEntries())
	settingsTab = container.NewTabItem("Settings", tabs.OpenSettings(w, refreshTab))

	tabsContainer := container.NewAppTabs(newTab, viewTab, settingsTab)
	tabsContainer.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabsContainer)
}
