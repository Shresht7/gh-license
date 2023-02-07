package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Shresht7/gh-license/api"
	"github.com/Shresht7/gh-license/helpers"
)

// ============
// LIST COMMAND
// ============

// List all licenses
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show a list of licenses",
	Long:  `Show a list of licenses available on GitHub.`,
	Example: helpers.ListExamples([]string{
		"gh-license list",
		"gh-license list --json",
		"gh-license list --pretty-json",
	}),
	Run: func(cmd *cobra.Command, args []string) {

		// Get list of licenses
		licenses, err := api.ListLicenses()
		if err != nil {
			fmt.Println(err)
			return
		}

		if cmd.Flag("web").Value.String() == "true" { // Check if web flag is set

			// Open the license in the browser
			err := helpers.OpenInBrowser("https://choosealicense.com/licenses/")
			if err != nil {
				fmt.Println(err)
				return
			}

		} else if cmd.Flag("pretty-json").Value.String() == "true" { // Check if pretty JSON flag is set

			// Prettify the JSON and print it
			output, err := helpers.Prettify(licenses)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(output)

		} else if cmd.Flag("json").Value.String() == "true" { // Check if JSON flag is set

			// Print the list of licenses in JSON format
			fmt.Println(licenses)

		} else { // Print the list of licenses in a table format by default

			// Print the list of licenses
			for _, license := range licenses {
				fmt.Printf("%-16s %s\n", license.Spdx_id, license.Name)
			}

		}

	},
}

// ----
// INIT
// ----

func init() {
	// Add list command
	rootCmd.AddCommand(listCmd)

	// Add Flags
	listCmd.Flags().BoolP("json", "j", false, "Output in JSON format")
	listCmd.Flags().BoolP("pretty-json", "p", false, "Output in pretty JSON format")
	listCmd.Flags().BoolP("web", "w", false, "Open the license in the browser")
}
