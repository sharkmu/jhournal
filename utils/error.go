package utils

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

var ShowErrorDialog = true

func DisplayError(err error) {
	if ShowErrorDialog {
		var w fyne.Window
		windows := fyne.CurrentApp().Driver().AllWindows()
		if len(windows) > 0 {
			w = windows[0]
		}
		dialog.ShowError(err, w)
	}
}
