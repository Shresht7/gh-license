package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Shresht7/Scribe/helpers"
	md "github.com/Shresht7/Scribe/markdown"
	"github.com/Shresht7/sliceutils/slice"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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

	// Generate the markdown for the command and write it to the file
	md, err := generateMarkdown(cmd)
	if err != nil {
		return err
	}
	w.WriteString(md)

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

	// If all goes as planned, return nil (no error)
	return nil
}

// Generates markdown documentation for the given cobra command.
func generateMarkdown(cmd *cobra.Command) (string, error) {

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
			strings.Join(slice.Map(cmd.Aliases, func(alias string, i int) string {
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
	// Write the links
	doc.WriteUnorderedList(links)

	// Write the buffer to the writer
	// _, err := md.buf.WriteTo(w)
	// if err != nil {
	// 	return err
	// }

	return doc.String(), nil

}

// HELPER FUNCTIONS
// ----------------

func linkFile(name string) string {
	return md.Link(name, "./"+name+".md").String()
}
