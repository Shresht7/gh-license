package main

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Shresht7/Scribe/helpers"
	md "github.com/Shresht7/Scribe/markdown"
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
		"docs/templates/_command-full.md",
		"docs/templates/_back-to-top.md",
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

	i, err := os.Create(filepath.Join(dir, "README.md"))
	if err != nil {
		return err
	}
	defer i.Close()

	// Generate the index file
	md, err := generateIndex(cmd)
	if err != nil {
		return err
	}
	i.WriteString(md)

	// If all goes as planned, return nil (no error)
	return nil
}

func generateIndex(cmd *cobra.Command) (string, error) {

	// Create a new markdown document
	doc := md.NewDocument()

	// Write the title
	doc.WriteHeading(1, helpers.Wrap("`", cmd.Name()))

	// Write the short description
	doc.WriteParagraph(cmd.Short)

	// Write the long description
	if cmd.Long != "" {
		doc.WriteHeading(2, "Description")
		doc.WriteParagraph(cmd.Long)
	}

	// Write the usage
	doc.WriteHeading(2, "Usage")
	doc.WriteCodeBlock(strings.TrimRight(cmd.Example, "\n"))

	// Write the commands
	doc.WriteHeading(2, "Commands")

	commands := []string{}
	for _, subCmd := range cmd.Commands() {
		text := linkFile(subCmd.Name())
		if subCmd.Short != "" {
			text += " - " + subCmd.Short
		}
		commands = append(commands, text)
	}

	doc.WriteUnorderedList(commands)

	// Return the markdown
	return doc.String(), nil

}
