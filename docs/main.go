package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Shresht7/gh-license/cmd"
	"github.com/spf13/cobra"
)

// Generate generates the documentation for the given package.
func main() {

	// Generate documentation in the current directory
	err := GenerateDocumentation(cmd.RootCmd, "docs", true)
	if err != nil {
		log.Fatal(err)
	}

}

// GenerateMarkdown generates the markdown for the given command.
func CreateMarkdownFile(cmd *cobra.Command, path string) error {

	// create the markdown file
	w, err := os.Create(path)
	if err != nil {
		return err
	}
	defer w.Close()

	// generate markdown for the command
	err = GenerateMarkdown(cmd, w)
	if err != nil {
		return err
	}

	return nil
}

func GenerateDocumentation(cmd *cobra.Command, dir string, recurse bool) error {

	// Generate the markdown file for the command
	// Join the directory and the command name
	path := filepath.Join(dir, cmd.Name()+".md")
	err := CreateMarkdownFile(cmd, path)
	if err != nil {
		return err
	}

	// Recurse through the subcommands
	if recurse {
		for _, subCmd := range cmd.Commands() {

			// Generate the markdown file for the subcommand
			err := GenerateDocumentation(subCmd, dir, recurse)
			if err != nil {
				return err
			}

		}
	}

	return nil
}
