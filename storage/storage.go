package storage

import (
	"os"
	"path/filepath"
)

const (
	appDir   = ".tasks"
	dataFile = "tasks.csv"
)

func getDataDir() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	dataDirectory := filepath.Join(homeDir, appDir)
	return dataDirectory, nil
}

func GetDataFilePath() (string, error) {
	dataDir, err := getDataDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dataDir, dataFile), nil
}
