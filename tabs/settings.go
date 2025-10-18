package tabs

import (
	"log"
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
	pathEntry.SetPlaceHolder("Nothing selected")

	browseBtn := widget.NewButton("Browse", func() {
		fd := dialog.NewFolderOpen(func(r fyne.ListableURI, err error) {
			if r == nil {
				return
			}
			SaveToEnv(r.Path())
			pathEntry.SetText(r.Path())
		}, w)
		fd.Show()
	})

	pathBox := container.NewVBox(pathLabel, pathEntry, browseBtn)

	return container.NewVBox(title, pathBox)
}

func SaveToEnv(folderpath string) {
	env := map[string]string{
		"JSON_FOLDER_PATH": folderpath,
	}

	configPath, err := utils.GetConfigDir()
	if err != nil {
		log.Fatal("Unable to get config directory:", err)
	}

	err = godotenv.Write(env, filepath.Join(configPath, ".env"))
	if err != nil {
		log.Fatal("Unable to write .env: ", err)
	}
}
