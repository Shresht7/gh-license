package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Shresht7/gh-license/api"
	"github.com/Shresht7/gh-license/helpers"
)

// ============
// VIEW COMMAND
// ============

// View details about a particular license
var viewCmd = &cobra.Command{
	Use:     "view",
	Aliases: []string{"show", "get"},
	Short:   "View details about a particular license",
	Long:    `View details about a particular license. For a list of available licenses, use the 'list' command.`,
	Args:    cobra.ExactArgs(1),
	Example: helpers.ListExamples([]string{
		"gh-license view mit",
		"gh-license view mit --json",
		"gh-license view mit --pretty-json",
	}),
	Run: func(cmd *cobra.Command, args []string) {

		//	Get license name from args
		name := args[0]

		//	Get license details from GitHub
		license, err := api.GetLicense(name)
		if err != nil {
			fmt.Println(err)
			return
		}

		if cmd.Flag("web").Value.String() == "true" { // Check if web flag is set

			// Open the license in the browser
			err := helpers.OpenInBrowser(license.Html_url)
			if err != nil {
				fmt.Println(err)
				return
			}

		} else if cmd.Flag("pretty-json").Value.String() == "true" { // Check if pretty JSON flag is set

			// Prettify the JSON and print it
			output, err := helpers.Prettify(license)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(output)

		} else if cmd.Flag("json").Value.String() == "true" { // Check if JSON flag is set

			// Print the license details in JSON format
			fmt.Println(license)

		} else { // Print the license details in a table format by default

			//	Print information about the license
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

	},
}

// ----
// INIT
// ----

func init() {
	//	Add view command
	rootCmd.AddCommand(viewCmd)

	//	Add flags to view command
	viewCmd.Flags().BoolP("json", "j", false, "Print the license details in JSON format")
	viewCmd.Flags().BoolP("pretty-json", "p", false, "Print the license details in pretty JSON format")
	viewCmd.Flags().BoolP("web", "w", false, "Open the license in the browser")
}
