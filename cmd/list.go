package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Shresht7/gh-license/api"
)

//	============
//	LIST COMMAND
//	============

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List licenses",
	Long:  `Show a list of licenses`,
	Run: func(cmd *cobra.Command, args []string) {

		//	Get list of licenses
		licenses, err := api.ListLicenses()
		if err != nil {
			fmt.Println(err)
			return
		}

		//	Print the list of licenses
		printLicenses(licenses)

	},
}

// Print the list of licenses to the console
func printLicenses(licenses []api.LicenseSimple) {
	for _, license := range licenses {
		fmt.Printf("%-16s %s\n", license.Spdx_id, license.Name)
	}
}

func init() {
	//	Add list command
	rootCmd.AddCommand(listCmd)
}
