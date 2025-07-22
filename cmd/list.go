package cmd

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/segmentationfaulter/todo-cli-app/storage"
	"github.com/spf13/cobra"
)

var listAllTasks bool

var listCmd = cobra.Command{
	Use:   "list",
	Short: "Show list of pending tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := storage.ReadTasksFile()
		if err != nil {
			return err
		}

		if !listAllTasks {
			tasks = slices.DeleteFunc(tasks, func(task []string) bool {
				completed, _ := strconv.ParseBool(task[3])
				return completed
			})
		}

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)

		fmt.Fprintln(w, "ID\tTask\tCreated at\tDone\t")
		for _, task := range tasks {
			timestamp, err := formatTime(task[2])
			if err != nil {
				return err
			}
			fmt.Fprintf(w, "%v\t%v\t%v\t%v\t\n", task[0], task[1], timestamp, task[3])
		}
		w.Flush()

		return nil
	},
}

func formatTime(epochStr string) (string, error) {
	unixTime, err := strconv.ParseInt(epochStr, 10, 64)

	if err != nil {
		return "", err
	}

	tm := time.Unix(unixTime, 0)
	return tm.Format(time.RFC822), nil
}

func init() {
	listCmd.Flags().BoolVarP(&listAllTasks, "all", "a", listAllTasks, "Show both pending and completed tasks")
	rootCmd.AddCommand(&listCmd)
}
