package utils

import (
	"errors"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func DisplayError(msg string) {
	var w fyne.Window
	windows := fyne.CurrentApp().Driver().AllWindows()
	if len(windows) > 0 {
		w = windows[0]
	}

	err := errors.New(msg)
	dialog.ShowError(err, w)
}
