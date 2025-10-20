package tabs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/sharkmu/jhournal/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func OpenSettings(w fyne.Window, onSaved func(string)) fyne.CanvasObject {
	title := widget.NewLabelWithStyle("Settings", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	pathLabel := widget.NewLabel("JSON's folder location: ")

	pathEntry := widget.NewEntry()

	configPath, err := utils.GetConfigDir()
	if err != nil {
		utils.DisplayError(fmt.Sprintf("Unable to get config directory: %v", err))
	}
	envPath := filepath.Join(configPath, ".env")

	_, err = os.Stat(envPath)
	if os.IsNotExist(err) {
		pathEntry.SetPlaceHolder("Nothing selected")
	} else if err != nil {
		utils.DisplayError(fmt.Sprintf("Error checking .env file: %v", err))
	} else {
		err := godotenv.Load(envPath)
		if err != nil {
			utils.DisplayError(fmt.Sprintf("Error loading .env file: %v", err))
		}
		pathEntry.Text = os.Getenv("JSON_FOLDER_PATH")
	}

	saveBtn := widget.NewButton("Save", func() {
		utils.SaveJsonToEnv(pathEntry.Text)
		onSaved("view")
	})

	pathInput := container.NewBorder(nil, nil, nil, saveBtn, pathEntry)

	browseBtn := widget.NewButton("Browse", func() {
		fd := dialog.NewFolderOpen(func(r fyne.ListableURI, err error) {
			if r == nil {
				return
			}
			utils.SaveJsonToEnv(r.Path())
			pathEntry.SetText(r.Path())
			onSaved("view")
		}, w)
		fd.Show()
	})

	pathBox := container.NewVBox(pathLabel, pathInput, browseBtn)

	sizeLabel := widget.NewLabel("Window's size")

	sizeLabelH := widget.NewLabel("Height:")
	sizeEntryH := widget.NewEntry()
	sizeEntryHContainer := container.New(
		layout.NewGridWrapLayout(fyne.NewSize(50, sizeEntryH.MinSize().Height)),
		sizeEntryH,
	)

	sizeLabelW := widget.NewLabel("Width:")
	sizeEntryW := widget.NewEntry()
	sizeEntryWContainer := container.New(
		layout.NewGridWrapLayout(fyne.NewSize(50, sizeEntryW.MinSize().Height)),
		sizeEntryW,
	)

	sizeEmptySpace := widget.NewLabel("")

	sizeSaveBtn := widget.NewButton("Save", func() {
		utils.SaveSizeToEnv(sizeEntryH.Text, sizeEntryW.Text)
	})
	sizeEntries := container.NewHBox(sizeLabelH, sizeEntryHContainer, sizeLabelW, sizeEntryWContainer, sizeEmptySpace, sizeSaveBtn)

	sizeBox := container.NewVBox(sizeLabel, sizeEntries)

	return container.NewVBox(title, pathBox, sizeBox)
}
