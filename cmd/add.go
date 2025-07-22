package cmd

import (
	"strconv"
	"time"

	"github.com/segmentationfaulter/todo-cli-app/storage"
	"github.com/spf13/cobra"
)

var addCmd = cobra.Command{
	Use:   "add",
	Short: "Add a new task to our todo list",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, tasks, err := nextId()
		if err != nil {
			return err
		}

		task := []string{id, args[0], strconv.FormatInt(time.Now().Unix(), 10), strconv.FormatBool(false)}
		tasks = append(tasks, task)

		err = SaveTasksFile(tasks)
		if err != nil {
			return err
		}

		return nil
	},
}

func nextId() (string, [][]string, error) {
	tasks, err := storage.ReadTasksFile()
	if err != nil {
		return "", nil, err
	}
	tasksCount := len(tasks)

	if tasksCount == 0 {
		return "1", [][]string{}, nil
	}

	nextId, err := strconv.Atoi(tasks[tasksCount-1][0])
	if err != nil {
		return "", nil, err
	}

	return strconv.Itoa(nextId + 1), tasks, nil
}

func init() {
	rootCmd.AddCommand(&addCmd)
}
