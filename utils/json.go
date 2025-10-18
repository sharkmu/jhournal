package utils

import (
	"encoding/json"
	"log"
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

func WriteJson(d []Data, jsonPath string, text string) {
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

func LenJson() int {
	d, _ := ReadJson()
	return len(d)
}
