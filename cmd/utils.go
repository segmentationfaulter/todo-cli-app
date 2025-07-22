package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/segmentationfaulter/todo-cli-app/storage"
)

func validateArgAsInteger(args []string) error {
	_, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%q is not a valid task ID", args[0])
	}
	return nil
}

func updateTasks(cmdArgs []string, updateFn func([][]string) [][]string) error {
	err := validateArgAsInteger(cmdArgs)
	if err != nil {
		return err
	}

	tasks, err := storage.ReadTasksFile()
	if err != nil {
		return err
	}

	tasks = updateFn(tasks)
	err = SaveTasksFile(tasks)
	if err != nil {
		return err
	}
	return nil
}

func SaveTasksFile(tasks [][]string) error {
	dataDir, err := storage.GetDataDir()

	if err != nil {
		return err
	}

	tmpFilePath := filepath.Join(dataDir, "tmp")

	tmpFile, err := os.OpenFile(tmpFilePath, os.O_RDWR|os.O_CREATE, 0600)

	if err != nil {
		return err
	}

	if err := syscall.Flock(int(tmpFile.Fd()), syscall.LOCK_EX); err != nil {
		tmpFile.Close()
		return err
	}

	writer := csv.NewWriter(tmpFile)
	err = writer.WriteAll(tasks)
	if err != nil {
		storage.CloseFile(tmpFile)
		return err
	}

	dataFilePath, err := storage.GetDataFilePath()
	if err != nil {
		return err
	}

	err = os.Rename(tmpFilePath, dataFilePath)
	if err != nil {
		storage.CloseFile(tmpFile)
		return err
	}

	return storage.CloseFile(tmpFile)
}
