package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Shresht7/gh-license/api"
	"github.com/Shresht7/gh-license/helpers"
)

//	============
//	REPO COMMAND
//	============

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:     "repo",
	Aliases: []string{"r"},
	Short:   "View license of a repository",
	Long:    `View license of a repository. Please provide the repository name in the format 'owner/repo'.`,
	Example: `gh license repo Shresht7/gh-license`,
	Run: func(cmd *cobra.Command, args []string) {

		//	Get owner and repo name
		owner, repo, err := helpers.DetermineOwnerAndRepo(args)
		if err != nil {
			fmt.Println(err)
			return
		}

		//	Get license details
		license, err := api.GetRepoLicense(owner, repo)
		if err != nil {
			fmt.Println(err)
			return
		}

		//	Print information about the license
		fmt.Println(license.Name)

	},
}

func init() {
	//	Add repo command
	rootCmd.AddCommand(repoCmd)
}
