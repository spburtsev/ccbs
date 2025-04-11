package cmd

import (
	"fmt"

	"github.com/spburtsev/ccbs/bootstrapping"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a project in the current directory",
	Run: func(cmd *cobra.Command, args []string) {
		err := bootstrapping.ExecInit()
		if err != nil {
			fmt.Printf("Could not initialize a project in current directory:\n%s\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
