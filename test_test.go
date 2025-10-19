package main

import (
	"testing"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/test"
	"github.com/sharkmu/jhournal/tabs"
	"github.com/sharkmu/jhournal/utils"
)

func TestGetConfigDir(t *testing.T) {
	dir, err := utils.GetConfigDir()
	if err != nil {
		t.Fatalf("Unable to get config dir %v", dir)
	}
	if dir == "" {
		t.Fatalf("Config dir is empty %v", dir)
	}
}

func TestTabs(t *testing.T) {
	w := test.NewWindow(nil)
	defer w.Close()

	var newTab *container.TabItem
	var viewTab *container.TabItem
	var settingsTab *container.TabItem

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
			t.Fatalf("No such tab to refresh %s", tabName)
		}
	}

	newTab = container.NewTabItem("New Entry", tabs.NewEntry(refreshTab))

	viewTab = container.NewTabItem("View Entries", tabs.ViewEntries())

	settingsTab = container.NewTabItem("Settings", tabs.OpenSettings(w, refreshTab))

	tabsContainer := container.NewAppTabs(newTab, viewTab, settingsTab)
	tabsContainer.SetTabLocation(container.TabLocationLeading)

	if newTab == nil || viewTab == nil || settingsTab == nil {
		t.Fatal("One or more tabs are nil")
	}

	w.SetContent(tabsContainer)

	if tabsContainer == nil {
		t.Fatal("AppTabs creation failed")
	}
}
