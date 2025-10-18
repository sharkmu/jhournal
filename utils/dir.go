package utils

import (
	"log"
	"os"
	"path/filepath"
)

func GetConfigDir() (string, error) {
	configRoot, err := os.UserConfigDir()
	if err != nil {
		log.Fatal("Unable to get user config directory:", err)
	}

	configDir := filepath.Join(configRoot, "jhournal")
	os.MkdirAll(configDir, 0755)
	return configDir, nil
}
