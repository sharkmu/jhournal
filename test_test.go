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

	var viewTab *container.TabItem

	refreshView := func() {
		viewTab.Content = tabs.ViewEntries()
	}

	newTab := container.NewTabItem("New Entry", tabs.NewEntry(refreshView))

	viewTab = container.NewTabItem("View Entries", tabs.ViewEntries())

	settingsTab := container.NewTabItem("Settings", tabs.OpenSettings(w))

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
