package tabs

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/sharkmu/jhournal/utils"
)

func ViewEntries() fyne.CanvasObject {
	var content fyne.CanvasObject
	jsonLength, err := utils.LenJson()
	if err != nil {
		utils.DisplayError(fmt.Errorf("unable to get the length of the JSON file: %w", err))
	}
	if jsonLength > 0 {
		entries, _, err := utils.ReadJson()
		if err != nil {
			utils.DisplayError(fmt.Errorf("unable to read JSON file: %w", err))
		}

		display := widget.NewLabel("Choose an entry from the list")

		listWidget := widget.NewList(
			func() int { return jsonLength },
			func() fyne.CanvasObject { return widget.NewLabel("") },
			func(i widget.ListItemID, o fyne.CanvasObject) {
				d := entries[i]
				o.(*widget.Label).SetText(fmt.Sprintf("%d. — %s", d.Id, d.Time.Format("2006-01-02 15:04")))
			},
		)

		listWidget.OnSelected = func(id widget.ListItemID) {
			d := entries[id]

			display.SetText(fmt.Sprintf(
				"ID: %d \nTime: %s\n\n------\n\n%s", d.Id,
				d.Time.Format("Monday, 2006-01-02 15:04:05"), d.Content),
			)
			display.Wrapping = fyne.TextWrapWord
		}

		split := container.NewHSplit(listWidget, display)
		split.SetOffset(0.3)
		content = split
	} else {
		label := widget.NewLabelWithStyle(
			"No entries yet. Head over to New Entry to make the first entry.",
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		)
		content = container.NewCenter(label)
	}

	return content
}
