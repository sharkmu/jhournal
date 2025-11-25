package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func SaveEnvValues(updates map[string]string) error {
	configPath, err := GetConfigDir()
	envPath := filepath.Join(configPath, ".env")

	if err != nil {
		return fmt.Errorf("unable to get config directory: %w", err)
	}

	_, err = os.Stat(envPath)
	var existing map[string]string
	if os.IsNotExist(err) {
		err := os.WriteFile(envPath, []byte(""), 0644)
		if err != nil {
			return fmt.Errorf("failed to create .env: %w", err)
		}
		existing = make(map[string]string)
	} else if err != nil {
		return fmt.Errorf("unable to get data on .env: %w", err)
	} else {
		existing, err = godotenv.Read(envPath)
		if err != nil {
			return fmt.Errorf("can't read existing .env file: %w", err)
		}
	}

	for k, v := range updates {
		existing[k] = v
	}

	err = godotenv.Write(existing, envPath)
	if err != nil {
		return fmt.Errorf("unable to write .env: %w", err)
	}

	return nil
}

func SaveJsonToEnv(folderpath string) error {
	err := SaveEnvValues(map[string]string{"JSON_FOLDER_PATH": folderpath})
	if err != nil {
		return err
	}
	return nil
}

func SaveWidthToEnv(sizeW string) error {
	err := SaveEnvValues(map[string]string{"WINDOW_SIZE_W": sizeW})
	if err != nil {
		return err
	}
	return nil
}

func SaveHeightToEnv(sizeH string) error {
	err := SaveEnvValues(map[string]string{"WINDOW_SIZE_H": sizeH})
	if err != nil {
		return err
	}
	return nil
}
