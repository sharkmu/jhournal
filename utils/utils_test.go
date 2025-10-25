package utils_test

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/sharkmu/jhournal/utils"
)

func TestGetConfigDir(t *testing.T) {
	configDir, err := utils.GetConfigDir()
	if err != nil {
		t.Fatalf("Unable to get config dir: %v", err)
	}
	if configDir == "" {
		t.Fatal("Config directory should not be empty")
	}

	info, err := os.Stat(configDir)
	if err != nil {
		t.Fatalf("Config directory doesn't exist: %v", err)
	}

	if !info.IsDir() {
		t.Fatal("Config path is not a directory")
	}

	if filepath.Base(configDir) != "jhournal" {
		t.Errorf("Expected config directory to end with 'jhournal', got: %s", configDir)
	}
}

func TestSaveJsonToEnv(t *testing.T) {
	configDir, err := utils.GetConfigDir()
	if err != nil {
		t.Fatalf("Unable to get config dir: %v", err)
	}

	testPath := filepath.Join(configDir, "test_folder")
	err = utils.SaveJsonToEnv(testPath)
	if err != nil {
		t.Fatalf("SaveJsonToEnv failed: %v", err)
	}

	envPath := filepath.Join(configDir, ".env")
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		t.Fatal(".env file was not created")
	}
}

func TestSaveSizeToEnv(t *testing.T) {
	configDir, err := utils.GetConfigDir()
	if err != nil {
		t.Fatalf("Unable to get config dir: %v", err)
	}

	testWidth := "800"
	testHeight := "600"

	utils.SaveWidthToEnv(testWidth)
	utils.SaveHeightToEnv(testHeight)

	envPath := filepath.Join(configDir, ".env")
	if _, err := os.Stat(envPath); os.IsNotExist(err) {
		t.Fatal(".env file was not created")
	}
}

func TestReadJson(t *testing.T) {
	configDir, err := utils.GetConfigDir()
	if err != nil {
		t.Fatalf("Unable to get config dir: %v", err)
	}

	testDir := filepath.Join(configDir, "test_data")
	err = os.MkdirAll(testDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
	defer os.RemoveAll(testDir)

	err = utils.SaveJsonToEnv(testDir)
	if err != nil {
		t.Fatalf("Failed to save json path to env: %v", err)
	}

	jsonPath := filepath.Join(testDir, "data.json")
	testData := []utils.Data{
		{Id: 1, Content: "Test entry 1", Time: time.Now()},
		{Id: 2, Content: "Test entry 2", Time: time.Now()},
	}

	jsonBytes, err := json.MarshalIndent(testData, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal test data: %v", err)
	}

	err = os.WriteFile(jsonPath, jsonBytes, 0644)
	if err != nil {
		t.Fatalf("Failed to write test JSON file: %v", err)
	}

	data, returnedPath := utils.ReadJson()

	if len(data) != 2 {
		t.Errorf("Expected 2 entries, got %d", len(data))
	}

	if data[0].Content != "Test entry 1" {
		t.Errorf("Expected 'Test entry 1', got '%s'", data[0].Content)
	}

	if returnedPath != jsonPath {
		t.Errorf("Expected path '%s', got '%s'", jsonPath, returnedPath)
	}
}

func TestWriteJson(t *testing.T) {
	configDir, err := utils.GetConfigDir()
	if err != nil {
		t.Fatalf("Unable to get config dir: %v", err)
	}

	testDir := filepath.Join(configDir, "test_write")
	err = os.MkdirAll(testDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
	defer os.RemoveAll(testDir)

	jsonPath := filepath.Join(testDir, "data.json")

	initialData := []utils.Data{
		{Id: 1, Content: "First entry", Time: time.Now()},
	}

	newText := "New test entry"
	utils.WriteJson(initialData, jsonPath, newText)

	fileData, err := os.ReadFile(jsonPath)
	if err != nil {
		t.Fatalf("Failed to read written JSON file: %v", err)
	}

	var resultData []utils.Data
	err = json.Unmarshal(fileData, &resultData)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	if len(resultData) != 2 {
		t.Errorf("Expected 2 entries after write, got %d", len(resultData))
	}

	if resultData[1].Content != newText {
		t.Errorf("Expected content '%s', got '%s'", newText, resultData[1].Content)
	}

	if resultData[1].Id != 2 {
		t.Errorf("Expected ID 2, got %d", resultData[1].Id)
	}
}

func TestLenJson(t *testing.T) {
	configDir, err := utils.GetConfigDir()
	if err != nil {
		t.Fatalf("Unable to get config dir: %v", err)
	}

	testDir := filepath.Join(configDir, "test_len")
	err = os.MkdirAll(testDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}
	defer os.RemoveAll(testDir)

	err = utils.SaveJsonToEnv(testDir)
	if err != nil {
		t.Fatalf("Failed to save json path to env: %v", err)
	}

	jsonPath := filepath.Join(testDir, "data.json")
	testData := []utils.Data{
		{Id: 1, Content: "Entry 1", Time: time.Now()},
		{Id: 2, Content: "Entry 2", Time: time.Now()},
		{Id: 3, Content: "Entry 3", Time: time.Now()},
	}

	jsonBytes, err := json.MarshalIndent(testData, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal test data: %v", err)
	}

	err = os.WriteFile(jsonPath, jsonBytes, 0644)
	if err != nil {
		t.Fatalf("Failed to write test JSON file: %v", err)
	}

	length := utils.LenJson()

	if length != 3 {
		t.Errorf("Expected length 3, got %d", length)
	}
}
