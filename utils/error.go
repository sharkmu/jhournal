package utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func DisplayError(err error) error {
	var w fyne.Window
	windows := fyne.CurrentApp().Driver().AllWindows()
	if len(windows) > 0 {
		w = windows[0]
	}

	dialog.ShowError(err, w)
	return err
}
