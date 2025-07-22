package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

var completeCmd = cobra.Command{
	Use:   "complete",
	Short: "Mark a task as competed",
	Args:  cobra.MatchAll(cobra.ExactArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		return updateTasks(args, func(tasks [][]string) [][]string {
			for _, task := range tasks {
				if task[0] == args[0] {
					task[3] = strconv.FormatBool(true)
					break
				}
			}

			return tasks
		})
	},
}

func init() {
	rootCmd.AddCommand(&completeCmd)
}
