package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetConfigDir() (string, error) {
	configRoot, err := os.UserConfigDir()
	if err != nil {
		DisplayError(fmt.Errorf("unable to get user config directory: %w", err))
	}

	configDir := filepath.Join(configRoot, "jhournal")
	os.MkdirAll(configDir, 0755)
	return configDir, nil
}
