package cmd

import (
	"davidallendj/gitman/util"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var commandCmd = &cobra.Command{
	Use:   "command",
	Short: "Manage command aliases to execute",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			util.PrintMap(Config.Commands)
		}
	},
}

func init() {
	commandCmd.AddCommand(&cobra.Command{
		Use:   "add",
		Short: "Add a command alias",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			// Check for at least two args
			if len(args) < 2 {
				log.Errorf("could not add command\n")
				return
			}

			fmt.Printf("%v: %v", args[0], args[1])

			// Add a new command and update config
			var commands map[string]string = viper.GetStringMapString("commands")
			commands[args[0]] = args[1]
			viper.Set("commands", commands)
			viper.WriteConfig()
		},
	}, &cobra.Command{
		Use:   "remove",
		Short: "Remove a command alias",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			// Check for command alias
			if len(args) < 1 {
				log.Errorf("could not remove command alias")
				return
			}

			fmt.Printf("%v\n", args[0])

			// Remove command aliase and update config
			var commands map[string]string = viper.GetStringMapString("commands")
			delete(commands, args[0])
			viper.Set("commands", commands)
			viper.WriteConfig()
		},
	})
	rootCmd.AddCommand(commandCmd)
}
