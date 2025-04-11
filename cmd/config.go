package cmd

import (
	"fmt"

	"github.com/spburtsev/ccbs/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Perform actions with ccbs configuration",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var configListSubcommand = &cobra.Command{
	Use:   "list",
	Short: "Display the contents of the config file",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.ReadGlobalConfig()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(cfg)
		}
	},
}

var configResetSubcommand = &cobra.Command{
	Use:   "reset",
	Short: "Reset the global configuration to default values",
	Run: func(cmd *cobra.Command, args []string) {
		err := config.ResetGlobalConfig()
		if err != nil {
			fmt.Printf("Error while resetting: %s\n", err)
			return
		}
		fmt.Println("Global config successfully reset to defaults")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.AddCommand(configListSubcommand)
	configCmd.AddCommand(configResetSubcommand)
}
