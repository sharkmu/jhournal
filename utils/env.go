package utils

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

func SaveToEnv(folderpath string) {
	env := map[string]string{
		"JSON_FOLDER_PATH": folderpath,
	}

	configPath, err := GetConfigDir()
	if err != nil {
		log.Fatal("Unable to get config directory:", err)
	}

	err = godotenv.Write(env, filepath.Join(configPath, ".env"))
	if err != nil {
		log.Fatal("Unable to write .env: ", err)
	}
}
