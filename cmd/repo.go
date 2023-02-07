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
	}),
	Run: func(cmd *cobra.Command, args []string) {

		// Get owner and repo name from args
		owner, repo, err := helpers.DetermineOwnerAndRepo(args)
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

	},
}

// ----
// INIT
// ----

func init() {
	//	Add repo command
	rootCmd.AddCommand(repoCmd)
}
