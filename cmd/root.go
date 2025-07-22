package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = (&cobra.Command{
	Use:     "tasks",
	Short:   "A todo app which runs in terminal",
	Version: "1.0.0",
})

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
