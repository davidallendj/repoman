package cmd

import (
	"davidallendj/gitman/util"
	"fmt"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a command alias",
	Long:  "Run a command alias for a collection of repositories",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Errorf("could not run command alias\n")
			return
		}

		// Run a command alias on all groups passed
		for _, command := range args {
			for _, group := range Config.Repositories {
				run := Config.Commands[command]
				cs := strings.Split(run, " ")
				c := exec.Command(cs[0], cs[1:]...)
				c.Dir = group
				stdout, err := c.Output()
				if err != nil {
					log.Errorf("could not run process: %v\n", err)
				}
				fmt.Println(string(stdout))
			}
		}
	},
}

func init() {
	runCmd.Flags().StringArrayVar(&Repositories, "repos", util.GetKeys(Config.Repositories), "Set groups of repositories")
	rootCmd.AddCommand(runCmd)
}
