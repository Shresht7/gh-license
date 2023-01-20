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

//	FLAGS
//	-----

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

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a license file",
	Long:  `Create a license file for your project`,
	Run: func(cmd *cobra.Command, args []string) {

		// Get license name
		name := args[0]

		// Get license details
		license, err := api.GetLicense(name)
		if err != nil {
			fmt.Println(err)
			return
		}

		//	Fill in placeholders
		contents := helpers.FillInPlaceholders(license.Body, map[string]string{
			"author":      author,
			"year":        year,
			"project":     project,
			"description": description,
		})

		//	Determine destination
		var dest string
		if len(output) > 0 {
			dest = output
		} else {
			dest = "LICENSE"
		}

		//	Write license file
		os.WriteFile(dest, []byte(contents), 0644)

	},
}

func init() {
	//	Add create command
	rootCmd.AddCommand(createCmd)

	//	Flags
	//	-----

	createCmd.Flags().StringVarP(&author, "author", "a", "", "Author of the project")
	createCmd.Flags().StringVarP(&year, "year", "y", strconv.Itoa(time.Now().Year()), "Year")
	createCmd.Flags().StringVarP(&project, "project", "p", "", "Project name")
	createCmd.Flags().StringVarP(&description, "description", "d", "", "Project description")
	createCmd.Flags().StringVarP(&output, "output", "o", "LICENSE", "Filepath")
}
