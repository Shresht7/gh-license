package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Shresht7/gh-license/api"
)

//	============
//	VIEW COMMAND
//	============

var viewCmd = &cobra.Command{
	Use:     "view",
	Short:   "View details about a particular license",
	Long:    `View details about a particular license. For a list of available licenses, use the 'list' command.`,
	Args:    cobra.ExactArgs(1),
	Example: `gh license view mit`,
	Run: func(cmd *cobra.Command, args []string) {

		//	Get license name
		name := args[0]

		//	Get license details
		license, err := api.GetLicense(name)
		if err != nil {
			fmt.Println(err)
			return
		}

		//	Print information about the license
		printDetails(license)

	},
}

// Print details about the license to the console
func printDetails(license api.License) {
	fmt.Println("\nName:", license.Name)
	fmt.Println("SPDX_ID:", license.Spdx_id)
	fmt.Println("URL:", license.Html_url)
	fmt.Println("\nDescription:\n\n", license.Description)
	fmt.Println("\nImplementation:\n\n", license.Implementation)
	fmt.Println("\nPermissions:", license.Permissions)
	fmt.Println("Conditions:", license.Conditions)
	fmt.Println("Limitations:", license.Limitations)
	fmt.Println("\nBody:\n\n", license.Body)
}

func init() {
	//	Add view command
	rootCmd.AddCommand(viewCmd)
}
