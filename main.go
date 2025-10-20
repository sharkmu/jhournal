package main

import (
	"log"
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

func createWindow(wSize fyne.Size) {
	a := app.NewWithID("com.sharkmu.jhournal")
	w := a.NewWindow("Jhournal")
	w.Resize(wSize)
	w.CenterOnScreen()

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
			log.Fatal("No such tab to refresh", tabName)
		}
	}

	newTab = container.NewTabItem("New Entry", tabs.NewEntry(refreshTab))

	viewTab = container.NewTabItem("View Entries", tabs.ViewEntries())

	settingsTab = container.NewTabItem("Settings", tabs.OpenSettings(w, refreshTab))

	tabsContainer := container.NewAppTabs(newTab, viewTab, settingsTab)
	tabsContainer.SetTabLocation(container.TabLocationLeading)

	w.SetContent(tabsContainer)
	w.ShowAndRun()
}

func getSize() fyne.Size {
	configPath, err := utils.GetConfigDir()
	if err != nil {
		log.Fatal("Unable to get config directory:", err)
	}
	envPath := filepath.Join(configPath, ".env")

	_, err = os.Stat(envPath)
	if os.IsNotExist(err) {
		utils.SaveJsonToEnv(configPath)
	}
	err = godotenv.Overload(envPath)
	if err != nil {
		log.Fatal("Error loading .env file:", err)
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
		log.Fatal(err)
	}

	sizeWf64, err := strconv.ParseFloat(sizeW, 32)
	if err != nil {
		log.Fatal(err)
	}

	return fyne.NewSize(float32(sizeWf64), float32(sizeHf64))
}

func main() {
	wSize := getSize()
	createWindow(wSize)
}
