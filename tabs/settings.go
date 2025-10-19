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

func OpenSettings(w fyne.Window, onSaved func(string)) fyne.CanvasObject {
	title := widget.NewLabelWithStyle("Settings", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	pathLabel := widget.NewLabel("JSON's folder location: ")

	pathEntry := widget.NewEntry()

	configPath, err := utils.GetConfigDir()
	if err != nil {
		log.Fatal("Unable to get config directory:", err)
	}
	envPath := filepath.Join(configPath, ".env")

	_, err = os.Stat(envPath)
	if os.IsNotExist(err) {
		pathEntry.SetPlaceHolder("Nothing selected")
	} else if err != nil {
		log.Fatal("Error checking .env file:", err)
	} else {
		err := godotenv.Load(envPath)
		if err != nil {
			log.Fatal("Error loading .env file:", err)
		}
		pathEntry.Text = os.Getenv("JSON_FOLDER_PATH")
	}

	saveBtn := widget.NewButton("Save", func() {
		utils.SaveToEnv(pathEntry.Text)
		onSaved("view")
	})

	pathInput := container.NewBorder(nil, nil, nil, saveBtn, pathEntry)

	browseBtn := widget.NewButton("Browse", func() {
		fd := dialog.NewFolderOpen(func(r fyne.ListableURI, err error) {
			if r == nil {
				return
			}
			utils.SaveToEnv(r.Path())
			pathEntry.SetText(r.Path())
			onSaved("view")
		}, w)
		fd.Show()
	})

	pathBox := container.NewVBox(pathLabel, pathInput, browseBtn)

	return container.NewVBox(title, pathBox)
}
