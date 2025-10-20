package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func SaveEnvValues(updates map[string]string) {
	configPath, err := GetConfigDir()
	envPath := filepath.Join(configPath, ".env")

	if err != nil {
		DisplayError(fmt.Sprintf("Unable to get config directory: %v", err))
	}

	_, err = os.Stat(envPath)
	var existing map[string]string
	if os.IsNotExist(err) {
		err := os.WriteFile(envPath, []byte(""), 0644)
		if err != nil {
			DisplayError(fmt.Sprintf("Failed to create .env: %v", err))
		}
		existing = make(map[string]string)
	} else if err != nil {
		DisplayError(fmt.Sprintf("Unable to get data on .env: %v", err))
	} else {
		existing, err = godotenv.Read(envPath)
		if err != nil {
			DisplayError(fmt.Sprintf("Can't read existing .env file: %v", err))
		}
	}

	for k, v := range updates {
		existing[k] = v
	}

	err = godotenv.Write(existing, envPath)
	if err != nil {
		DisplayError(fmt.Sprintf("Unable to write .env: %v", err))
	}
}

func SaveJsonToEnv(folderpath string) error {
	SaveEnvValues(map[string]string{"JSON_FOLDER_PATH": folderpath})
	return nil
}

func SaveSizeToEnv(sizeH, sizeW string) {
	SaveEnvValues(map[string]string{
		"WINDOW_SIZE_H": sizeH,
		"WINDOW_SIZE_W": sizeW,
	})
}
