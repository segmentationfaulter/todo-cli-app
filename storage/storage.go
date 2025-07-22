package storage

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"syscall"
)

const (
	appDir   = ".tasks"
	dataFile = "tasks.csv"
)

func GetDataDir() (string, error) {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	dataDirectory := filepath.Join(homeDir, appDir)
	if err := os.MkdirAll(dataDirectory, 0755); err != nil {
		return "", err
	}
	return dataDirectory, nil
}

func GetDataFilePath() (string, error) {
	dataDir, err := GetDataDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dataDir, dataFile), nil
}

func OpenTasksFile() (*os.File, error) {
	path, err := GetDataFilePath()

	if err != nil {
		return nil, err
	}

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		return nil, err
	}

	if err := syscall.Flock(int(file.Fd()), syscall.LOCK_EX); err != nil {
		file.Close()
		return nil, err
	}

	return file, nil
}

func CloseFile(file *os.File) error {
	syscall.Flock(int(file.Fd()), syscall.LOCK_UN)
	return file.Close()
}

func ReadTasksFile() ([][]string, error) {
	file, err := OpenTasksFile()

	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)
	tasks, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}
	err = CloseFile(file)

	return tasks, err
}
