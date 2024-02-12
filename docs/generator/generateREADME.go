package main

import (
	"html/template"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// Generates the main README.md file for the project.
func generateREADME(cmd *cobra.Command) error {

	// Create the README.md file
	w, err := os.Create("README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	// Create and parse the template definitions
	tmpl := template.Must(template.ParseFiles(
		"docs/generator/templates/README.md",
		"docs/generator/templates/_back-to-top.md",
		"docs/generator/templates/_command-readme.md",
	))

	// Execute the template and write the README.md file
	err = tmpl.Execute(w, toTemplateData(cmd))
	if err != nil {
		return err
	}

	return nil
}
