package main

import (
	"log"
	"strings"

	"github.com/Shresht7/Scribe/helpers"
	"github.com/Shresht7/Scribe/markdown"
	"github.com/Shresht7/sliceutils/slice"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/Shresht7/gh-license/cmd"
)

// Generate generates the documentation for this project
func main() {

	// Generate the documentation for the root command and all subcommands
	// The documentation will be generated in the docs directory
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

// HELPER FUNCTIONS
// ----------------

// toTemplateData converts a cobra command to a map of strings that can be used
// in a template. This is used to generate the documentation for a command.
func toTemplateData(cmd *cobra.Command) map[string]string {
	return map[string]string{

		"Name": cmd.Name(),

		"Description": cmd.Short,

		"LongDescription": cmd.Long,

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
}
