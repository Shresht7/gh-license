package main

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/Shresht7/Scribe/markdown"
	"github.com/Shresht7/goutils/slice"
	str "github.com/Shresht7/goutils/strings"
)

// HELPER FUNCTIONS
// ----------------

// toTemplateData converts a cobra command to a map of strings that can be used
// in a template. This is used to generate the documentation for a command.
func toTemplateData(cmd *cobra.Command) map[string]any {
	return map[string]any{

		"Name": cmd.Name(),

		"Description": cmd.Short,

		"LongDescription": cmd.Long,

		"Aliases": strings.Join(
			slice.Map(cmd.Aliases, func(x string, i int) string {
				return str.Wrap("`", x)
			}),
			", ",
		),

		"Usage": cmd.UseLine(),

		"Commands": (func() []map[string]any {
			commands := []map[string]any{}
			for _, subCmd := range cmd.Commands() {
				commands = append(commands, toTemplateData(subCmd))
			}
			return commands
		})(),

		"Flags": (func() string {
			flags := cmd.Flags()
			if flags == nil || !flags.HasFlags() {
				return ""
			}

			// Instantiate the headers
			headers := []string{"Flag", "Type", "Description", "Default"}

			// Instantiate the rows
			rows := [][]string{}
			cmd.Flags().VisitAll(func(flag *pflag.Flag) {

				// Make the default value generic
				switch flag.Name {
				case "author":
					flag.DefValue = "[AuthorName]"
				case "project":
					flag.DefValue = "[RepositoryName]"
				case "year":
					flag.DefValue = "[CurrentYear]"
				}

				// Append the row
				rows = append(rows, []string{
					"`--" + flag.Name + ", -" + flag.Shorthand + "`",
					str.Wrap("`", flag.Value.Type()),
					flag.Usage,
					flag.DefValue,
				})
			})

			return markdown.Table(headers, rows)
		})(),

		"Examples": strings.TrimRight(cmd.Example, "\n"),

		"SeeAlso": (func() []string {
			links := []string{}
			// Link to the parent command
			if cmd.HasParent() {
				links = append(links, linkFile(cmd.Parent().Name()))

				// Link to the other sibling commands
				for _, subCmd := range cmd.Parent().Commands() {
					if subCmd != cmd {
						links = append(links, linkFile(subCmd.Name()))
					}
				}
			}
			// Link to the child commands
			for _, subCmd := range cmd.Commands() {
				links = append(links, linkFile(subCmd.Name()))
			}

			return links
		})(),
	}
}

// linkFile returns a markdown link to the given file name.
func linkFile(name string) string {
	return markdown.Link(name, "./"+name+".md")
}
