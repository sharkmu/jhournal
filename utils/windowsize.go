package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"fyne.io/fyne/v2"
	"github.com/joho/godotenv"
	"github.com/kbinani/screenshot"
)

func GetSize() fyne.Size {
	configPath, err := GetConfigDir()
	if err != nil {
		DisplayError(fmt.Errorf("unable to get config directory: %w", err))
	}
	envPath := filepath.Join(configPath, ".env")

	_, err = os.Stat(envPath)
	if os.IsNotExist(err) {
		SaveJsonToEnv(configPath)
	}
	err = godotenv.Overload(envPath)
	if err != nil {
		DisplayError(fmt.Errorf("error loading .env file: %w", err))
	}

	sizeW := os.Getenv("WINDOW_SIZE_W")
	sizeH := os.Getenv("WINDOW_SIZE_H")
	if sizeW == "" {
		sizeW = "800"
	}
	if sizeH == "" {
		sizeH = "600"
	}

	sizeWf64, err := strconv.ParseFloat(sizeW, 32)
	if err != nil {
		DisplayError(fmt.Errorf("error: %w", err))
	}

	sizeHf64, err := strconv.ParseFloat(sizeH, 32)
	if err != nil {
		DisplayError(fmt.Errorf("error: %w", err))
	}

	return fyne.NewSize(float32(sizeWf64), float32(sizeHf64))
}

func GetScreenSize() fyne.Size {
	bounds := screenshot.GetDisplayBounds(0)
	width := bounds.Dx() - 150
	height := bounds.Dy() - 150

	return fyne.NewSize(float32(width), float32(height))
}
