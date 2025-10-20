package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"github.com/joho/godotenv"
	"github.com/sharkmu/jhournal/tabs"
	"github.com/sharkmu/jhournal/utils"
)

func main() {
	a := app.NewWithID("com.sharkmu.jhournal")
	w := a.NewWindow("Jhournal")

	w.CenterOnScreen()

	w.Show()

	wSize := getSize()
	w.Resize(wSize)

	createTabs(w)

	a.Run()
}

func getSize() fyne.Size {
	configPath, err := utils.GetConfigDir()
	if err != nil {
		utils.DisplayError(fmt.Errorf("unable to get config directory: %w", err))
	}
	envPath := filepath.Join(configPath, ".env")

	_, err = os.Stat(envPath)
	if os.IsNotExist(err) {
		utils.SaveJsonToEnv(configPath)
	}
	err = godotenv.Overload(envPath)
	if err != nil {
		utils.DisplayError(fmt.Errorf("error loading .env file: %w", err))
	}

	sizeH := os.Getenv("WINDOW_SIZE_H")
	sizeW := os.Getenv("WINDOW_SIZE_W")
	if sizeH == "" {
		sizeH = "600"
	}
	if sizeW == "" {
		sizeW = "800"
	}

	sizeHf64, err := strconv.ParseFloat(sizeH, 32)
	if err != nil {
		utils.DisplayError(fmt.Errorf("error: %w", err))
	}

	sizeWf64, err := strconv.ParseFloat(sizeW, 32)
	if err != nil {
		utils.DisplayError(fmt.Errorf("error: %w", err))
	}

	return fyne.NewSize(float32(sizeWf64), float32(sizeHf64))
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
