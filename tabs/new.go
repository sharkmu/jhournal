package tabs

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/sharkmu/jhournal/utils"
)

type Data struct {
	Id      int64
	Content string
	Time    time.Time
}

func NewEntry(onSaved func(tabName string)) fyne.CanvasObject {
	title := widget.NewLabelWithStyle("New Entry", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	warning1 := canvas.NewText("", color.RGBA{R: 255, G: 0, B: 0, A: 255})
	warning2 := canvas.NewText("", color.RGBA{R: 255, G: 0, B: 0, A: 255})
	warning1Status := false
	warning2Status := false

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Write here")

	clearWarnings := func() {
		if warning1Status {
			warning1.Text = ""
			warning1.Refresh()
			warning1Status = false
		}
		if warning2Status {
			warning2.Text = ""
			warning2.Refresh()
			warning2Status = false
		}
	}

	button := widget.NewButton("Save", func() {
		if entry.Text != "" {
			d, jsonPath, err := utils.ReadJson()
			fmt.Println(jsonPath)
			if err != nil {
				utils.DisplayError(fmt.Errorf("unable to read JSON file: %v", err))
			}

			handleSavingFunc := func() {
				err = utils.WriteJson(d, jsonPath, entry.Text)
				if err != nil {
					utils.DisplayError(err)
				}
				clearWarnings()
				entry.Text = ""
				entry.Refresh()
				onSaved("view")
			}

			if len(d) == 0 {
				handleSavingFunc()
			} else {
				lastDate := d[len(d)-1].Time
				sinceLastDate := time.Since(lastDate)

				if sinceLastDate >= time.Hour {
					handleSavingFunc()
				} else {
					elapsed := sinceLastDate
					mins := int(elapsed.Minutes())
					secs := int(elapsed.Seconds()) % 60

					warning2.Text = fmt.Sprintf(
						"60 minutes hasn't passed yet. Time since last entry: %02dm %02ds",
						mins, secs,
					)
					warning2.Refresh()
					warning2Status = true
					return
				}
			}
		}
	})

	maxChars := 77 // If increased by a lot, consider changing entry to widget.NewMultiLineEntry()

	entry.OnChanged = func(s string) {
		if len(s) >= maxChars {
			entry.SetText(s[:maxChars])
			warning1.Text = fmt.Sprintf("Max characters (%d) reached!", maxChars)
			warning1.Refresh()
			warning1Status = true
		} else {
			if warning1Status {
				warning1.Text = ""
				warning1.Refresh()
				warning1Status = false
			}
		}

		if entry.Text == "" {
			if warning2Status {
				warning2.Text = ""
				warning2.Refresh()
				warning2Status = false
			}
		}
	}

	return container.NewVBox(title, warning1, warning2, entry, button)
}
