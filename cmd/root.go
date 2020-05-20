package cmd

import (
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "goignore",
		Short: "A CLI tool for generating .gitignore files",
		Long: ` --------------------------------------------
| A CLI tool for generating .gitignore files |
 --------------------------------------------
Usage: goignore COMMAND [OPTIONS]
Help:  goignore help`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func exit(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile,
		"config",
		"",
		"config file (default is $HOME/.goignore.yaml)",
	)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			exit(err)
		}
		viper.AddConfigPath(home)
		viper.SetConfigFile(".goignore")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}