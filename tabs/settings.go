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

	pathLabel := widget.NewLabel("JSON location: ")

	pathEntry := widget.NewEntry()
	pathEntry.SetPlaceHolder("Nothing selected")

	browseBtn := widget.NewButton("Browse", func() {
		fd := dialog.NewFileOpen(func(r fyne.URIReadCloser, err error) {
			if r == nil {
				return
			}
			defer r.Close()
			saveToEnv(r.URI().Path())
			pathEntry.SetText(r.URI().Path())
		}, w)
		fd.Show()
	})

	pathBox := container.NewVBox(pathLabel, pathEntry, browseBtn)

	return container.NewVBox(title, pathBox)
}

func saveToEnv(filepath string) {
	env := map[string]string{
		"JSON_FILE_PATH": filepath,
	}

	err := godotenv.Write(env, "./.env")
	if err != nil {
		log.Fatal("Unable to write .env: ", err)
	}
}
