package cmd

import (
	"fmt"
	"strings"

	"github.com/cli/go-gh"
	"github.com/spf13/cobra"

	"github.com/Shresht7/gh-license/api"
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
		var owner, repo string
		if len(args) == 0 { // If no arguments are provided, get the current repository
			current, err := gh.CurrentRepository()
			if err != nil {
				fmt.Println(err)
				return
			}
			owner, repo = current.Owner(), current.Name()
		} else { // If arguments are provided, split them into owner and repo
			splitArgs := strings.Split(args[0], "/")
			owner, repo = splitArgs[0], splitArgs[1]
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
