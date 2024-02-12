package main

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

// Generates the documentation for the given cobra command. If recurse is true,
// it will recurse through the subcommands and generate documentation for them as well
func generateDocumentation(cmd *cobra.Command, dir string, recurse bool) error {

	// Generate the markdown file for the command
	path := filepath.Join(dir, cmd.Name()+".md")
	w, err := os.Create(path)
	if err != nil {
		return err
	}
	defer w.Close()

	// Create and parse the template definitions
	tmpl := template.Must(template.ParseFiles(
		"docs/generator/templates/command.md",
		"docs/generator/templates/_back-to-top.md",
	))

	err = tmpl.Execute(w, toTemplateData(cmd))
	if err != nil {
		return err
	}

	// Recurse through the subcommands...
	if recurse {
		for _, subCmd := range cmd.Commands() {

			// ... and generate documentation for them
			err := generateDocumentation(subCmd, dir, recurse)
			if err != nil {
				return err
			}

		}
	}

	// Create the index file for all the commands
	i, err := os.Create(filepath.Join(dir, "README.md"))
	if err != nil {
		return err
	}
	defer i.Close()

	// Create and parse the template definitions
	tmpl = template.Must(template.ParseFiles(
		"docs/generator/templates/index.md",
		"docs/generator/templates/_back-to-top.md",
	))

	err = tmpl.Execute(i, toTemplateData(cmd))
	if err != nil {
		return err
	}

	// If all goes as planned, return nil (no error)
	return nil
}
