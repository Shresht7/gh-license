/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/cli/go-gh"
	"github.com/spf13/cobra"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View details about a particular license",
	Long:  `View details about a particular license`,
	Run: func(cmd *cobra.Command, args []string) {

		//	Get gh client
		client, err := gh.RESTClient(nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		license := "licenses/" + args[0]

		//	Fetch the license
		response := struct{ Body string }{}
		err = client.Get(license, &response)
		if err != nil {
			fmt.Println(err)
			return
		}

		//	Print the list of licenses
		fmt.Println(response)

	},
}

//	TODO: Add --repo flag to target specific repo

func init() {
	rootCmd.AddCommand(viewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
