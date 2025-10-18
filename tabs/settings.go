package tabs

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/sharkmu/jhournal/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func OpenSettings(w fyne.Window) fyne.CanvasObject {
	title := widget.NewLabelWithStyle("Settings", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	pathLabel := widget.NewLabel("JSON's folder location: ")

	pathEntry := widget.NewEntry()

	configPath, err := utils.GetConfigDir()
	if err != nil {
		log.Fatal("Unable to get config directory:", err)
	}
	envPath := filepath.Join(configPath, ".env")
	var placeholderText string

	_, err = os.Stat(envPath)
	if os.IsNotExist(err) {
		placeholderText = "Nothing selected"
	} else if err != nil {
		log.Fatal("Error checking .env file:", err)
	} else {
		err := godotenv.Load(envPath)
		if err != nil {
			log.Fatal("Error loading .env file:", err)
		}
		placeholderText = os.Getenv("JSON_FOLDER_PATH")
	}
	pathEntry.SetPlaceHolder(placeholderText)

	browseBtn := widget.NewButton("Browse", func() {
		fd := dialog.NewFolderOpen(func(r fyne.ListableURI, err error) {
			if r == nil {
				return
			}
			utils.SaveToEnv(r.Path())
			pathEntry.SetText(r.Path())
		}, w)
		fd.Show()
	})

	pathBox := container.NewVBox(pathLabel, pathEntry, browseBtn)

	return container.NewVBox(title, pathBox)
}
