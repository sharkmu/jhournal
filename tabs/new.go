package tabs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func NewEntry() fyne.CanvasObject {
	return widget.NewLabel("New entry")
}
