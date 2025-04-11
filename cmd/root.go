package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ccbs",
	Short: "ccbs is a cli tool for bootstrapping cmake projects",
	Long:  "ccbs is a cli tool for bootstrapping cmake projects",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing ccbs '%s'\n", err)
		os.Exit(1)
	}
}
