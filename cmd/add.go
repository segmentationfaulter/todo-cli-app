package cmd

import (
	"encoding/csv"
	"os"

	"github.com/segmentationfaulter/todo-cli-app/storage"
	"github.com/spf13/cobra"
)

var addCmd = cobra.Command{
	Use:   "add",
	Short: "Add a new task to our todo list",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		path, err := storage.GetDataFilePath()

		if err != nil {
			return err
		}

		tasksFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			return err
		}
		writer := csv.NewWriter(tasksFile)

		if err := writer.Write(args); err != nil {
			return err
		}

		writer.Flush()
		if err := writer.Error(); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(&addCmd)
}
