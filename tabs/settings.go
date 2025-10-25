package tabs

import (
	"fmt"
	"image/color"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sharkmu/jhournal/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
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
		utils.DisplayError(fmt.Errorf("unable to get config directory: %w", err))
	}
	envPath := filepath.Join(configPath, ".env")

	_, err = os.Stat(envPath)
	if os.IsNotExist(err) {
		pathEntry.SetPlaceHolder("Nothing selected")
	} else if err != nil {
		utils.DisplayError(fmt.Errorf("error checking .env file: %w", err))
	} else {
		err := godotenv.Load(envPath)
		if err != nil {
			utils.DisplayError(fmt.Errorf("error loading .env file: %w", err))
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

	warningTextWidth := canvas.NewText("", color.RGBA{R: 255, G: 0, B: 0, A: 255})
	warningTextHeight := canvas.NewText("", color.RGBA{R: 255, G: 0, B: 0, A: 255})

	sizeLabelW := widget.NewLabel("Width:")
	sizeEntryW := widget.NewEntry()
	sizeEntryWContainer := container.New(
		layout.NewGridWrapLayout(fyne.NewSize(50, sizeEntryW.MinSize().Height)),
		sizeEntryW,
	)

	sizeLabelH := widget.NewLabel("Height:")
	sizeEntryH := widget.NewEntry()
	sizeEntryHContainer := container.New(
		layout.NewGridWrapLayout(fyne.NewSize(50, sizeEntryH.MinSize().Height)),
		sizeEntryH,
	)

	sizeSaveBtn := widget.NewButton("Save", func() {
		warningStatusW := false
		warningStatusH := false

		sizeWf64, err := strconv.ParseFloat(sizeEntryW.Text, 32)
		if err != nil {
			if sizeEntryW.Text == "" {
				sizeWf64 = 0
			} else {
				utils.DisplayError(fmt.Errorf("error: %w", err))
			}
		}

		sizeHf64, err := strconv.ParseFloat(sizeEntryH.Text, 32)
		if err != nil {
			if sizeEntryH.Text == "" {
				sizeHf64 = 0
			} else {
				utils.DisplayError(fmt.Errorf("error: %w", err))
			}
		}

		clearWarning := func(option string) {
			switch option {
			case "width":
				warningTextWidth.Text = ""
				warningTextWidth.Refresh()
			case "height":
				warningTextHeight.Text = ""
				warningTextHeight.Refresh()
			}
		}

		screenWidth := float64(utils.GetScreenSize().Width)
		screenHeight := float64(utils.GetScreenSize().Height)

		if sizeWf64 < 300 || sizeWf64 > screenWidth {
			warningTextWidth.Text = fmt.Sprintf(
				"  Incorrect value for width (min: 300, max: %v)",
				screenWidth,
			)
			warningTextWidth.Refresh()
			warningStatusW = true
		}

		if sizeHf64 < 220 || sizeHf64 > screenHeight {
			warningTextHeight.Text = fmt.Sprintf(
				"  Incorrect value for height (min: 220, max: %v)",
				screenHeight,
			)
			warningTextHeight.Refresh()
			warningStatusH = true
		}

		if !warningStatusW {
			utils.SaveWidthToEnv(strconv.FormatFloat(sizeWf64, 'f', -1, 64))
			clearWarning("width")
		}
		if !warningStatusH {
			utils.SaveHeightToEnv(strconv.FormatFloat(sizeHf64, 'f', -1, 64))
			clearWarning("height")
		}

	})
	sizeEntries := container.NewHBox(sizeLabelW, sizeEntryWContainer, sizeLabelH, sizeEntryHContainer, sizeSaveBtn)

	sizeBox := container.NewVBox(sizeLabel, sizeEntries, warningTextWidth, warningTextHeight)

	return container.NewVBox(title, pathBox, sizeBox)
}
