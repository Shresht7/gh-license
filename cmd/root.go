package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/Shresht7/gh-license/helpers"
)

//	============
//	ROOT COMMAND
//	============

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "gh-license",
	Version: "0.1.0",
	Short:   "List and create license files",
	Long:    `List and create license files using GitHub's license API`,
	Example: helpers.ListExamples([]string{
		"gh license list",
		"gh license view mit",
		"gh license create mit",
		"gh license create mit --author \"John Doe\"",
		"gh license create mit --year 2020 --author \"John Doe\"",
	}),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gh-license.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
