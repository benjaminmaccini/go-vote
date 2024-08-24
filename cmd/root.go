package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	. "git.sr.ht/~bmaccini/go-vote/pkg/utils"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use: "go-vote",
	Long: `A CLI app that handles elections.

Ex.
go-vote serve --candidates alice,bob --port 1337 --protocol simpleMajority
go-vote vote --electionId uuid --url 127.0.0.1 --port 1337 --ballot '{"alice": 1}'

or set the values with a config file (default $HOME/.go-vote.yaml)
`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		Logger.Fatal("", err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-vote.yaml)")
	rootCmd.PersistentFlags().StringP("port", "", "1337", "port that the election is hosted on")
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			Logger.Fatal("", err)
		}

		// Search config in home directory with name ".go-vote" (without extension).
		viper.AddConfigPath(home + "/.config/go-vote")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		Logger.Info("Config successfully loaded", "fileName", viper.ConfigFileUsed())
	}
}
