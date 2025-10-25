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
		DisplayError(fmt.Errorf("unable to get config directory: %w", err))
	}

	_, err = os.Stat(envPath)
	var existing map[string]string
	if os.IsNotExist(err) {
		err := os.WriteFile(envPath, []byte(""), 0644)
		if err != nil {
			DisplayError(fmt.Errorf("failed to create .env: %w", err))
		}
		existing = make(map[string]string)
	} else if err != nil {
		DisplayError(fmt.Errorf("unable to get data on .env: %w", err))
	} else {
		existing, err = godotenv.Read(envPath)
		if err != nil {
			DisplayError(fmt.Errorf("can't read existing .env file: %w", err))
		}
	}

	for k, v := range updates {
		existing[k] = v
	}

	err = godotenv.Write(existing, envPath)
	if err != nil {
		DisplayError(fmt.Errorf("unable to write .env: %w", err))
	}
}

func SaveJsonToEnv(folderpath string) error {
	SaveEnvValues(map[string]string{"JSON_FOLDER_PATH": folderpath})
	return nil
}

func SaveWidthToEnv(sizeW string) {
	SaveEnvValues(map[string]string{
		"WINDOW_SIZE_W": sizeW,
	})
}

func SaveHeightToEnv(sizeH string) {
	SaveEnvValues(map[string]string{
		"WINDOW_SIZE_H": sizeH,
	})
}
