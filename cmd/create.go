package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"

	"github.com/Shresht7/gh-license/api"
	"github.com/Shresht7/gh-license/helpers"
)

//	==============
//	CREATE COMMAND
//	==============

// Create a license file
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new", "add", "init", "set"},
	Short:   "Create a license file",
	Long:    `Create a license file for your project.`,
	Example: helpers.ListExamples([]string{
		"gh-license create mit",
		"gh-license create mit --author Shresht7 --year 2023",
		"gh-license create --web",
		"gh-license create mit --web",
	}),
	Run: func(cmd *cobra.Command, args []string) {

		// License Name
		var name string
		if len(args) > 0 {
			name = args[0]
		}

		// Get flags
		author := cmd.Flag("author").Value.String()
		repo := cmd.Flag("project").Value.String()
		year := cmd.Flag("year").Value.String()
		project := cmd.Flag("project").Value.String()
		description := cmd.Flag("description").Value.String()
		output := cmd.Flag("output").Value.String()
		shouldLaunchWeb := cmd.Flag("web").Value.String() == "true"

		var err error // Declare error variable

		if shouldLaunchWeb { // Check if web flag is set

			// Open the license in the browser
			url := determineUrl(name, author, repo, output)
			err = helpers.OpenInBrowser(url)
			if err != nil {
				fmt.Println(err)
				return
			}

			return // Exit the function
		}

		// Check if license name is provided before proceeding
		if name == "" {
			fmt.Println("Please provide a license name. Run `gh license list` to see a list of available licenses.")
			return // Exit the function
		}

		// Get license details
		license, err := api.GetLicense(name)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Fill in placeholders in license body
		contents := helpers.FillInPlaceholders(license.Body, map[string]string{
			"author":      author,
			"year":        year,
			"project":     project,
			"description": description,
		})

		// Write license file to the destination
		os.WriteFile(output, []byte(contents), 0644)

	},
}

// Determine URL to create license file using the web interface
func determineUrl(name, author, repo, output string) string {
	// Determine URL to create license
	url := fmt.Sprintf("https://github.com/%s/%s/community/license/new", author, repo)

	// Add query params
	queryParams := []string{}
	if name != "" {
		queryParams = append(queryParams, "template="+name)
	}
	if output != "" {
		queryParams = append(queryParams, "filename="+output)
	}

	// Convert to query string and add to URL
	queryString := strings.Join(queryParams, "&")
	if queryString != "" {
		url += "?" + queryString
	}

	// Return URL
	return url
}

// ----
// INIT
// ----

func init() {
	// Add create command
	rootCmd.AddCommand(createCmd)

	// Determine owner and repo from current directory
	owner, repo, err := helpers.DetermineOwnerAndRepo()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Add flags to create command
	createCmd.Flags().StringP("author", "a", owner, "Author of the project")
	createCmd.Flags().StringP("year", "y", strconv.Itoa(time.Now().Year()), "Year")
	createCmd.Flags().StringP("project", "p", repo, "Project name")
	createCmd.Flags().StringP("description", "d", "", "Project description")
	createCmd.Flags().StringP("output", "o", "LICENSE", "Filepath")
	createCmd.Flags().BoolP("web", "w", false, "Create license file using the web interface")
}
