package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "waterflow",
	Short: "Waterflow is a task scheduler",
	Long:  "WaterFlow is a go based task scheduler and workflow managment tool that simplifies the task automation through YAML configuraiton and CLI commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("WaterFlow v1.0")
	},
}

func init() {
	rootCommand.AddCommand(createCommand)
	rootCommand.AddCommand(udpateCommand)
	rootCommand.AddCommand(listCommand)
	rootCommand.AddCommand(deleteCommand)
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
