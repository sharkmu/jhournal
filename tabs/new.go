package tabs

import (
	"encoding/json"
	"fmt"
	"image/color"
	"log"
	"os"
	"path/filepath"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/joho/godotenv"
	"github.com/sharkmu/jhournal/utils"
)

type Data struct {
	Id      int64
	Content string
	Time    time.Time
}

func NewEntry() fyne.CanvasObject {
	title := widget.NewLabelWithStyle("New Entry", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	warning1 := canvas.NewText("", color.RGBA{R: 255, G: 0, B: 0, A: 255})
	warning2 := canvas.NewText("", color.RGBA{R: 255, G: 0, B: 0, A: 255})
	warning1Status := false
	warning2Status := false

	entry := widget.NewEntry()
	entry.SetPlaceHolder("Write here")

	clearWarnings := func() {
		entry.SetText("")
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
			d, jsonPath := readJson()
			if len(d) > 0 {
				lastDate := d[len(d)-1].Time
				sinceLastDate := time.Since(lastDate)

				if sinceLastDate >= time.Hour {
					writeJson(d, jsonPath, entry.Text)
					clearWarnings()
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
				}
			} else {
				writeJson(d, jsonPath, entry.Text)
				clearWarnings()
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

func readJson() ([]Data, string) {
	configPath, err := utils.GetConfigDir()
	if err != nil {
		log.Fatal("Unable to get config directory:", err)
	}
	envPath := filepath.Join(configPath, ".env")

	_, err = os.Stat(envPath)
	if os.IsNotExist(err) {
		SaveToEnv(configPath)
	} else if err != nil {
		log.Fatal("Error checking .env file:", err)
	} else {
		err := godotenv.Load(envPath)
		if err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	}

	jsonPath := filepath.Join(os.Getenv("JSON_FOLDER_PATH"), "data.json")

	fileData, err := os.ReadFile(jsonPath)
	if err != nil {
		if os.IsNotExist(err) {
			file, createErr := os.Create(jsonPath)
			if createErr != nil {
				log.Fatal("Error creating JSON file:", createErr)
			}
			defer file.Close()

			_, writeErr := file.Write([]byte("[]"))
			if writeErr != nil {
				log.Fatal("Error initialising JSON content:", writeErr)
			}

			fileData = []byte("[]")
		} else {
			log.Fatal("Unable to read JSON file:", err)
		}
	}

	var d []Data
	if len(fileData) > 0 {
		err = json.Unmarshal(fileData, &d)
		if err != nil {
			log.Fatal("Error parsing JSON:", err)
		}
	}
	return d, jsonPath
}

func writeJson(d []Data, jsonPath string, text string) {
	var index int64 = 1
	if len(d) > 0 {
		index = d[len(d)-1].Id + 1
	}

	newData := Data{index, text, time.Now()}
	d = append(d, newData)

	b, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		log.Fatal("Unable to marshal new data", err)
	}

	err = os.WriteFile(jsonPath, b, 0644)
	if err != nil {
		log.Fatal("Unable to write JSON file:", err)
	}
}
