package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
)

type Data struct {
	Id      int64
	Content string
	Time    time.Time
}

func ReadJson() ([]Data, string) {
	configPath, err := GetConfigDir()
	if err != nil {
		DisplayError(fmt.Sprintf("Unable to get config directory: %v", err))
	}
	envPath := filepath.Join(configPath, ".env")

	_, err = os.Stat(envPath)
	if os.IsNotExist(err) {
		SaveJsonToEnv(configPath)
	}
	err = godotenv.Overload(envPath)
	if err != nil {
		DisplayError(fmt.Sprintf("Error loading .env file: %v", err))
	}

	jsonPath := filepath.Join(os.Getenv("JSON_FOLDER_PATH"), "data.json")

	fileData, err := os.ReadFile(jsonPath)
	if err != nil {
		if os.IsNotExist(err) {
			file, createErr := os.Create(jsonPath)
			if createErr != nil {
				DisplayError(fmt.Sprintf("Error creating JSON file: %v", err))
			}
			defer file.Close()

			_, writeErr := file.Write([]byte("[]"))
			if writeErr != nil {
				DisplayError(fmt.Sprintf("Error initialising JSON content: %v", err))
			}

			fileData = []byte("[]")
		} else {
			DisplayError(fmt.Sprintf("Unable to read JSON file: %v", err))
		}
	}

	var d []Data
	if len(fileData) > 0 {
		err = json.Unmarshal(fileData, &d)
		if err != nil {
			DisplayError(fmt.Sprintf("Error parsing JSON: %v", err))
		}
	}
	return d, jsonPath
}

func WriteJson(d []Data, jsonPath string, text string) {
	var index int64 = 1
	if len(d) > 0 {
		index = d[len(d)-1].Id + 1
	}

	newData := Data{index, text, time.Now()}
	d = append(d, newData)

	b, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		DisplayError(fmt.Sprintf("Unable to marshal new data: %v", err))
	}

	err = os.WriteFile(jsonPath, b, 0644)
	if err != nil {
		DisplayError(fmt.Sprintf("Unable to write JSON file: %v", err))
	}
}

func LenJson() int {
	d, _ := ReadJson()
	return len(d)
}
