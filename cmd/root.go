package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper"
	"infracli/cmd/image"
	"infracli/cmd/project"
	"infracli/config"
	"os"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "infra",
	Short: "This is my infrastructure CLI",
	Long:  "Upload docker compose structure",
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

	rootCmd.AddCommand(project.Cmd)
	rootCmd.AddCommand(image.Cmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName("cobra")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}

	initRepositoryConfig()
}

func initRepositoryConfig() {
	name := viper.GetString("repository.name")

	repositoryConfig := config.RepositoryConfig{
		Name: name,
	}

	config.CreateRepository(repositoryConfig)
}
