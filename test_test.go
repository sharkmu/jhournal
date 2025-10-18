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

	newEntry := tabs.NewEntry()
	viewEntries := tabs.ViewEntries()
	settings := tabs.OpenSettings(w)

	if newEntry == nil || viewEntries == nil || settings == nil {
		t.Fatal("One or more tabs are nil")
	}

	tabsContent := container.NewAppTabs(
		container.NewTabItem("New Entry", newEntry),
		container.NewTabItem("View Entries", viewEntries),
		container.NewTabItem("Settings", settings),
	)

	if tabsContent == nil {
		t.Fatal("AppTabs creation failed")
	}
}
