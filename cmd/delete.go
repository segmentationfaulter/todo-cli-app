package cmd

import (
	"slices"

	"github.com/spf13/cobra"
)

var deleteCmd = cobra.Command{
	Use:   "delete",
	Short: "Delete a task by ID",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		return updateTasks(args, func(tasks [][]string) [][]string {
			return slices.DeleteFunc(tasks, func(task []string) bool {
				return args[0] == task[0]
			})
		})
	},
}

func init() {
	rootCmd.AddCommand(&deleteCmd)
}
