package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ConfigT struct {
	Repositories map[string]string `yaml:"repositories"`
	Commands     map[string]string `yaml:"commands"`
}

var (
	Config       ConfigT
	Repositories []string
)

var rootCmd = &cobra.Command{
	Use:   "gitman",
	Short: "Manage git repositories",
	Long:  "Manage git repositories more efficiently",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			c, _, _ := cmd.Find([]string{"help"})
			cmd.Run(c, []string{})
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(LoadConfig)
}

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.config/gitman")
	viper.AddConfigPath("/etc/gitman")
	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("could not find config file: %v", err)
		home := os.Getenv("HOME")
		err := os.MkdirAll(home+"/.config/gitman", 0o755)
		if err != nil {
			log.Errorf("could not create directory for config: %v", err)
		}
		viper.SetDefault("repositories", map[string]string{})
		viper.SetDefault("commands", map[string]string{})
		err = viper.WriteConfigAs(home + "/.config/gitman/config.yaml")
		if err != nil {
			log.Errorf("could not create default config: %v", err)
		}
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		} else {
			fmt.Println("Config file: ", viper.ConfigFileUsed())
		}
	}

	err = viper.Unmarshal(&Config)
	if err != nil {
		log.Errorf("could not load config: %v", err)
	}
}
