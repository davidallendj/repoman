package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Repositories map[string]string `yaml:"repositories"`
	Commands     map[string]string `yaml:"commands"`
}

var (
	config Config
)

var rootCmd = &cobra.Command{
	Use:   "gitman",
	Short: "Manage a collection of git repositories",
	Long:  "Created to make managing multiple repositories easier.",
	Run: func(cmd *cobra.Command, args []string) {

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
		err := os.MkdirAll("$HOME/.config/gitman", 0o755)
		if err != nil {
			log.Errorf("could not create directory for config: %v", err)
		}
		err = viper.WriteConfigAs("$HOME/.config/gitman/config.yaml")
		if err != nil {
			log.Errorf("could not create default config: %v", err)
		}
	}

	viper.SetDefault("repositories", map[string]string{})
	viper.SetDefault("commands", map[string]string{})
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		} else {
			fmt.Println("Config file: ", viper.ConfigFileUsed())
		}
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Errorf("could not load config: %v", err)
	}
}
