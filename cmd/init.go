package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a project in the current directory",
	Run: func(cmd *cobra.Command, args []string) {
		err := ExecInit()
		if err != nil {
			fmt.Printf("Could not initialize a project in current directory:\n%s\n", args[0], err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
