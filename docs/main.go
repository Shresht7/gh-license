package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Shresht7/Scribe/helpers"
	"github.com/Shresht7/Scribe/markdown"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/Shresht7/gh-license/cmd"
)

// Generate generates the documentation for the given package.
func main() {

	// Generate documentation in the current directory
	err := GenerateDocumentation(cmd.RootCmd, "docs", true)
	if err != nil {
		log.Fatal(err)
	}

	// Generate the README.md file
	err = GenerateREADME()
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
	md, err := GenerateMarkdown(cmd)
	if err != nil {
		return err
	}
	w.WriteString(md)

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

func GenerateREADME() error {

	// Generate the README.md file
	// Create the README.md file
	w, err := os.Create("README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	tmpl := template.Must(template.ParseFiles(
		"docs/templates/README.md",
		"docs/templates/_back-to-top.md",
		"docs/templates/_command.md",
	))

	err = tmpl.Execute(w, Map(cmd.RootCmd.Commands(), func(cmd *cobra.Command) map[string]string {
		return map[string]string{
			"Name":        cmd.Name(),
			"Description": cmd.Short,
			"Aliases": strings.Join(
				Map(cmd.Aliases, func(x string) string {
					return helpers.Wrap("`", x)
				}), ", "),
			"Usage": cmd.UseLine(),
			"Flags": (func() string {
				flags := cmd.Flags()
				if flags == nil || !flags.HasFlags() {
					return ""
				}

				headers := []string{"Flag", "Description", "Default"}
				rows := [][]string{}

				flags.VisitAll(func(flag *pflag.Flag) {
					rows = append(rows, []string{
						helpers.Wrap("`", flag.Name),
						flag.Usage,
						flag.DefValue,
					})
				})

				return markdown.Table(headers, rows).String()
			})(),
			"Examples": cmd.Example,
		}
	}))
	if err != nil {
		return err
	}

	return nil
}
