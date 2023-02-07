package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Shresht7/gh-license/api"
)

// ============
// LIST COMMAND
// ============

// List all licenses
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show a list of licenses",
	Long:  `Show a list of licenses available on GitHub.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Get list of licenses
		licenses, err := api.ListLicenses()
		if err != nil {
			fmt.Println(err)
			return
		}

		// Print the list of licenses
		for _, license := range licenses {
			fmt.Printf("%-16s %s\n", license.Spdx_id, license.Name)
		}

	},
}

// ----
// INIT
// ----

func init() {
	//	Add list command
	rootCmd.AddCommand(listCmd)
}
