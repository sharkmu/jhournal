package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetConfigDir() (string, error) {
	configRoot, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("unable to get user config directory: %w", err)
	}

	configDir := filepath.Join(configRoot, "jhournal")
	err = os.MkdirAll(configDir, 0755)
	if err != nil {
		return "", fmt.Errorf("unable to create config directory: %w", err)
	}
	return configDir, nil
}
