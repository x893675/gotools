package cmd

import (
	"config-parse/config"

	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:   "version",
	Short: "Get Versions of Clairctl and underlying services",
	Long:  `Get Versions of Clairctl and underlying services`,
	Run: func(cmd *cobra.Command, args []string) {
		config.Print()
	},
}

func init() {
	rootCmd.AddCommand(printCmd)
}
