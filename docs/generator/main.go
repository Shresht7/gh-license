package main

import (
	"log"

	"github.com/Shresht7/gh-license/cmd"
)

// Generates documentation for this project
func main() {

	// Generate the documentation for the root command and all subcommands in the "docs" folder
	err := generateDocumentation(cmd.RootCmd, "docs", true)
	if err != nil {
		log.Fatal(err)
	}

	// Generate the main README.md file
	err = generateREADME(cmd.RootCmd)
	if err != nil {
		log.Fatal(err)
	}

}
