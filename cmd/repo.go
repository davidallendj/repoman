package cmd

import (
	"davidallendj/gitman/util"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Verbose bool
)

var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Manage repositories",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			util.PrintMap(Config.Repositories)
		}
	},
}

func init() {
	repoCmd.AddCommand(&cobra.Command{
		Use:   "add",
		Short: "Add repo to be managed",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			// Check for at least two args (name and path)
			if len(args) < 2 {
				log.Errorf("could not add repository\n")
				return
			}

			fmt.Printf("%v: %v\n", args[0], args[1])

			// Add a new repo and update config
			var repos map[string]string = viper.GetStringMapString("repositories")
			repos[args[0]] = args[1]
			viper.Set("repositories", repos)
			viper.WriteConfig()
		},
	}, &cobra.Command{
		Use:   "remove",
		Short: "Remove repo from being managed",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			// Check for a single arg (name)
			if len(args) < 1 {
				log.Errorf("could not remove repository")
				return
			}

			fmt.Printf("%v\n", args[0])

			// Remove repository and update config
			var repos map[string]string = viper.GetStringMapString("repositories")
			delete(repos, args[0])
			viper.Set("repositories", repos)
			viper.WriteConfig()
		},
	})
	repoCmd.Flags().BoolVarP(&Verbose, "verbose", "v", false, "Set verbose output")
	rootCmd.AddCommand(repoCmd)
}
