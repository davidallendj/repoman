package cmd

import "github.com/spf13/cobra"

var commandCmd = &cobra.Command{
	Use:   "command",
	Short: "Add command alias to execute",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {

	rootCmd.AddCommand(commandCmd)
}
