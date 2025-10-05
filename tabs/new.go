package tabs

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewEntry() fyne.CanvasObject {
	title := widget.NewLabelWithStyle("New Entry", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	warning := canvas.NewText("", color.RGBA{R: 255, G: 0, B: 0, A: 255})
	warningStatus := false

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Write here")
	button := widget.NewButton("Save", func() {
		println("Saved new entry:", entry.Text)
		entry.SetText("")
		if warningStatus {
			warning.Text = ""
			warning.Refresh()
		}
	})

	maxChars := 77 // If increased by a lot, consider changing entry to widget.NewMultiLineEntry()

	entry.OnChanged = func(s string) {
		if len(s) >= maxChars {
			entry.SetText(s[:maxChars])
			warning.Text = fmt.Sprintf("Max characters (%d) reached!", maxChars)
			warning.Refresh()
			warningStatus = true
		} else {
			if warningStatus {
				warning.Text = ""
				warning.Refresh()
			}
		}
	}

	return container.NewVBox(title, warning, entry, button)
}
