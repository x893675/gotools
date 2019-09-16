package cmd

import (
	"config-parse/config"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "configparse",
	Short: "Analyze your docker image with Clair, directly from your registry or local images.",
	Long:  ``,
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	// rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
	// rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
	// rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	//rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	// viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	// viper.BindPFlag("projectbase", rootCmd.PersistentFlags().Lookup("projectbase"))
	// viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	// viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	// viper.SetDefault("license", "apache")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	// if err := config.Clean(); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(-1)
	// }
}

func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	config.Init(cfgFile)
}
