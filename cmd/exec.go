package cmd

import (
	"davidallendj/gitman/util"
	"fmt"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a command or script",
	Long:  "Execute a command or script instead of a command aliases",
	Run: func(cmd *cobra.Command, args []string) {
		// Execute any arbitrary command on a group
		if len(args) == 0 {
			log.Errorf("could not exec command\n")
			return
		}

		// Execute single command directly on provided groups
		for _, repo := range Config.Repositories {
			c := exec.Command(args[0], args[1:]...)
			c.Dir = repo
			stdout, err := c.Output()
			if err != nil {
				log.Errorf("could not run process: %v\n", err)
			}

			fmt.Println(string(stdout))
		}
	},
}

func init() {
	execCmd.Flags().StringArrayVar(&Repositories, "repos", util.GetKeys(Config.Repositories), "Set the groups of repositories")
	rootCmd.AddCommand(execCmd)
}
