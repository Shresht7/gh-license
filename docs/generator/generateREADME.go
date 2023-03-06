package main

import (
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/Shresht7/Scribe/helpers"
	"github.com/Shresht7/Scribe/markdown"
	"github.com/Shresht7/sliceutils/slice"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// Generates the main README.md file for the project.
// It uses templates from the docs/templates directory.
func generateREADME(cmd *cobra.Command) error {

	// Create the README.md file
	w, err := os.Create("README.md")
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	// Create and parse the template definitions
	tmpl := template.Must(template.ParseFiles(
		"docs/templates/README.md",
		"docs/templates/_back-to-top.md",
		"docs/templates/_command.md",
	))

	err = tmpl.Execute(w, slice.Map(cmd.Commands(), func(cmd *cobra.Command, idx int) map[string]string {
		return map[string]string{

			"Name": cmd.Name(),

			"Description": cmd.Short,

			"Aliases": strings.Join(
				slice.Map(cmd.Aliases, func(x string, i int) string {
					return helpers.Wrap("`", x)
				}),
				", ",
			),

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
