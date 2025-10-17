package tabs

import (
	"log"

	"github.com/joho/godotenv"

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
			saveToEnv(r.Path())
			pathEntry.SetText(r.Path())
		}, w)
		fd.Show()
	})

	pathBox := container.NewVBox(pathLabel, pathEntry, browseBtn)

	return container.NewVBox(title, pathBox)
}

func saveToEnv(filepath string) {
	env := map[string]string{
		"JSON_FOLDER_PATH": filepath,
	}

	err := godotenv.Write(env, "./.env")
	if err != nil {
		log.Fatal("Unable to write .env: ", err)
	}
}
