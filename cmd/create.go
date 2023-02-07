package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"

	"github.com/Shresht7/gh-license/api"
	"github.com/Shresht7/gh-license/helpers"
)

// -----
// FLAGS
// -----

var (
	author      string // Author of the project
	year        string // Year
	project     string // Project name
	description string // Project description
	output      string // Output filepath
)

//	==============
//	CREATE COMMAND
//	==============

// Create a license file
var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new", "add", "init", "set"},
	Args:    cobra.ExactArgs(1),
	Short:   "Create a license file",
	Long:    `Create a license file for your project.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Get license name from args
		name := args[0]

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

		// Determine destination of license file
		var dest string
		if len(output) > 0 {
			dest = output
		} else {
			dest = "LICENSE"
		}

		// Write license file to the destination
		os.WriteFile(dest, []byte(contents), 0644)

	},
}

// ----
// INIT
// ----

func init() {
	// Add create command
	rootCmd.AddCommand(createCmd)

	// Add flags to create command
	createCmd.Flags().StringVarP(&author, "author", "a", "", "Author of the project")
	createCmd.Flags().StringVarP(&year, "year", "y", strconv.Itoa(time.Now().Year()), "Year")
	createCmd.Flags().StringVarP(&project, "project", "p", "", "Project name")
	createCmd.Flags().StringVarP(&description, "description", "d", "", "Project description")
	createCmd.Flags().StringVarP(&output, "output", "o", "LICENSE", "Filepath")
}
