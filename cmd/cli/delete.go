package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCommand = &cobra.Command{
	Use:   "delete [workflow id]",
	Short: "Delete workflow from the scheduler",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configFile := args[0]
		fmt.Println("Creating a new workflow from ", configFile)
	},
}
