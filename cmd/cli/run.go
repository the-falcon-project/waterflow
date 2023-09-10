package cli

import (
	"github.com/spf13/cobra"
	"github.com/the-falcon-project/waterflow/cmd/app"
)

func runCommand() *cobra.Command {
	var detachMode string
	var runCommand = &cobra.Command{
		Use:   "run",
		Short: "starts the waterflow scheduler",
		Run: func(cmd *cobra.Command, args []string) {
			app.Run(detachMode)
		},
	}
	runCommand.Flags().StringVarP(&detachMode, "detach", "d", "false", "Detach mode allows WaterFlow to run in background of the terminal.")
	runCommand.Flags().Lookup("detach").NoOptDefVal = "true"
	return runCommand
}
