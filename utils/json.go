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

func ReadJson() ([]Data, string, error) {
	configPath, err := GetConfigDir()
	if err != nil {
		return nil, "", fmt.Errorf("unable to get config directory: %w", err)
	}
	envPath := filepath.Join(configPath, ".env")

	_, err = os.Stat(envPath)
	if os.IsNotExist(err) {
		err = SaveJsonToEnv(configPath)
		if err != nil {
			return nil, "", fmt.Errorf("failed to save JSON to env: %w", err)
		}
	}
	err = godotenv.Overload(envPath)
	if err != nil {
		return nil, "", fmt.Errorf("error loading .env file: %w", err)
	}

	jsonPath := filepath.Join(os.Getenv("JSON_FOLDER_PATH"), "data.json")

	fileData, err := os.ReadFile(jsonPath)
	if err != nil {
		if os.IsNotExist(err) {
			file, createErr := os.Create(jsonPath)
			if createErr != nil {
				return nil, "", fmt.Errorf("error creating JSON file: %w", err)
			}

			_, writeErr := file.Write([]byte("[]"))
			if writeErr != nil {
				return nil, "", fmt.Errorf("error initialising JSON content: %w", err)
			}
			err = file.Close()
			if err != nil {
				return nil, "", (fmt.Errorf("unable to close file: %w", err))
			}

			fileData = []byte("[]")
		} else {
			return nil, "", fmt.Errorf("unable to read JSON file: %w", err)
		}
	}

	var d []Data
	if len(fileData) > 0 {
		err = json.Unmarshal(fileData, &d)
		if err != nil {
			return nil, "", fmt.Errorf("error parsing JSON: %w", err)
		}
	}
	return d, jsonPath, nil
}

func WriteJson(d []Data, jsonPath string, text string) error {
	var index int64 = 1
	if len(d) > 0 {
		index = d[len(d)-1].Id + 1
	}

	newData := Data{index, text, time.Now()}
	d = append(d, newData)

	b, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to marshal new data: %w", err)
	}

	err = os.WriteFile(jsonPath, b, 0644)
	if err != nil {
		return fmt.Errorf("unable to write JSON file: %w", err)
	}

	return nil
}

func LenJson() (int, error) {
	d, _, err := ReadJson()
	if err != nil {
		return 0, fmt.Errorf("failed to read JSON: %w", err)
	}
	return len(d), nil
}
