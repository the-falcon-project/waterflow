package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var udpateCommand = &cobra.Command{
	Use:   "update [workflow id] [file]",
	Short: "Update existing workfow from a YAML configuration file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configFile := args[0]
		fmt.Println("Creating a new workflow from ", configFile)
	},
}
