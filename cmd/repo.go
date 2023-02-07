package cmd

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Shresht7/gh-license/api"
	"github.com/Shresht7/gh-license/helpers"
)

// ============
// REPO COMMAND
// ============

// View license of a repository
var repoCmd = &cobra.Command{
	Use:     "repo",
	Aliases: []string{"r"},
	Short:   "View license of a repository",
	Long:    `View license of a repository. If no repository is specified, the current repository is used.`,
	Example: helpers.ListExamples([]string{
		"gh license repo",
		"gh license repo Shresht7/gh-license",
		"gh license repo gh-license",
		"gh license repo --json",
		"gh license repo Shresht7/gh-license --pretty-json",
	}),
	Run: func(cmd *cobra.Command, args []string) {

		// Get owner and repo name from args
		owner, repo, err := helpers.DetermineOwnerAndRepo(args...)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Get license details
		license, err := api.GetRepoLicense(owner, repo)
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

			// Print the license in JSON format
			fmt.Println(license)

		} else { // Print the license in a table format by default

			// Print information about the license
			fmt.Println("License: " + license.License.Name)
			fmt.Println("URL: " + license.Html_url)
			fmt.Println("Description: " + license.Links.Self)
			content, err := base64.StdEncoding.DecodeString(license.Content)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Content:\n\n" + string(content))

		}

	},
}

// ----
// INIT
// ----

func init() {
	//	Add repo command
	rootCmd.AddCommand(repoCmd)

	// Add flags to repo command
	repoCmd.Flags().BoolP("json", "j", false, "Print the license in JSON format")
	repoCmd.Flags().BoolP("pretty-json", "p", false, "Print the license in pretty JSON format")
	repoCmd.Flags().BoolP("web", "w", false, "Open the license in the browser")
}
