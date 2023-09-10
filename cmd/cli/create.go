package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCommand = &cobra.Command{
	Use:   "create [file]",
	Short: "Create a new workflow from a YAML configuration file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configFile := args[0]
		fmt.Println("Creating a new workflow from ", configFile)
	},
}
