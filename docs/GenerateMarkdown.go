package main

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/Shresht7/Scribe/helpers"
	md "github.com/Shresht7/Scribe/markdown"
)

// GenerateMarkdown generates the markdown for the given command.
func GenerateMarkdown(cmd *cobra.Command) (string, error) {

	// Create a new markdown document
	doc := md.NewDocument()

	// Write the title
	doc.WriteHeading(1, helpers.Wrap("`", cmd.Name()))

	// Write the short Description
	doc.WriteParagraph(cmd.Short)

	// Write the aliases
	if len(cmd.Aliases) > 0 {
		doc.WriteHeading(2, "Aliases")
		doc.WriteParagraph(
			strings.Join(Map(cmd.Aliases, func(alias string) string {
				return helpers.Wrap("`", alias)
			}), ", "))
	}

	// Write the long description
	if cmd.Long != "" {
		doc.WriteHeading(2, "Description")
		doc.WriteParagraph(cmd.Long)
	}

	// Write the usage
	if cmd.Runnable() {
		doc.WriteHeading(2, "Usage")
		doc.WriteCodeBlock(cmd.UseLine())
	}

	// Write the flags
	if cmd.Flags().HasFlags() {
		doc.WriteHeading(2, "Flags")

		// Create a markdown table

		// Instantiate the headers
		headers := []string{"Flag", "Type", "Description", "Default"}

		// Instantiate the rows
		rows := [][]string{}
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			rows = append(rows, []string{
				"`--" + flag.Name + ", -" + flag.Shorthand + "`",
				helpers.Wrap("`", flag.Value.Type()),
				flag.Usage,
				flag.DefValue,
			})
		})

		// Write the table
		doc.WriteTable(headers, rows)
	}

	// Write the examples
	if len(cmd.Example) > 0 {
		doc.WriteHeading(2, "Examples")
		doc.WriteCodeBlock(cmd.Example)
	}

	// Write the see also section
	if cmd.HasParent() || len(cmd.Commands()) > 0 {
		doc.WriteHeading(2, "See Also")
	}

	links := []string{}
	// Link to the parent command
	if cmd.HasParent() {
		links = append(links, LinkFile(cmd.Parent().Name()))

		// Link to the other sibling commands
		for _, subCmd := range cmd.Parent().Commands() {
			if subCmd != cmd {
				links = append(links, LinkFile(subCmd.Name()))
			}
		}
	}
	// Link to the child commands
	for _, subCmd := range cmd.Commands() {
		links = append(links, LinkFile(subCmd.Name()))
	}
	// Write the links
	doc.WriteUnorderedList(links)

	// Write the buffer to the writer
	// _, err := md.buf.WriteTo(w)
	// if err != nil {
	// 	return err
	// }

	return doc.String(), nil

}

func LinkFile(name string) string {
	return md.Link(name, "./"+name+".md").String()
}
