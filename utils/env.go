package utils

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func SaveEnvValues(updates map[string]string) {
	configPath, err := GetConfigDir()
	envPath := filepath.Join(configPath, ".env")

	if err != nil {
		log.Fatal("Unable to get config directory:", err)
	}

	_, err = os.Stat(envPath)
	var existing map[string]string
	if os.IsNotExist(err) {
		err := os.WriteFile(envPath, []byte(""), 0644)
		if err != nil {
			log.Fatal("Failed to create .env:", err)
		}
		existing = make(map[string]string)
	} else if err != nil {
		log.Fatal("Unable to get data on .env:", err)
	} else {
		existing, err = godotenv.Read(envPath)
		if err != nil {
			log.Fatal("Can't read existing .env file:", err)
		}
	}

	for k, v := range updates {
		existing[k] = v
	}

	err = godotenv.Write(existing, envPath)
	if err != nil {
		log.Fatal("Unable to write .env:", err)
	}
}

func SaveJsonToEnv(folderpath string) {
	SaveEnvValues(map[string]string{"JSON_FOLDER_PATH": folderpath})
}

func SaveSizeToEnv(sizeH, sizeW string) {
	SaveEnvValues(map[string]string{
		"WINDOW_SIZE_H": sizeH,
		"WINDOW_SIZE_W": sizeW,
	})
}
