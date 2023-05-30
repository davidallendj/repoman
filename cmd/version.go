package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the current version",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Experimental git repository manager v0.0.1")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
