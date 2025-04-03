package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "yacb",
	Short: "yacb is a cli tool for bootstrapping cmake projects",
	Long:  "yacb is a cli tool for bootstrapping cmake projects",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing yacb '%s'\n", err)
		os.Exit(1)
	}
}
