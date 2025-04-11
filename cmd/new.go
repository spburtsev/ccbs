package cmd

import (
	"fmt"

	"github.com/spburtsev/ccbs/bootstrapping"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Args:  cobra.ExactArgs(1),
	Short: "Create a new project in specified path",
	Run: func(cmd *cobra.Command, args []string) {
		err := bootstrapping.ExecNew(args[0])
		if err != nil {
			fmt.Printf("Could not create a project in '%s':\n%s\n", args[0], err)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
