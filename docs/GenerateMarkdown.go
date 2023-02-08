package main

import (
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// GenerateMarkdown generates the markdown for the given command.
func GenerateMarkdown(cmd *cobra.Command, w io.Writer) error {

	// Create a new markdown writer
	md := NewMarkdownWriter()

	// Write the title
	md.WriteHeading(1, wrap("`", cmd.Name()))

	// Write the short Description
	md.WriteParagraph(cmd.Short)

	// Write the aliases
	if len(cmd.Aliases) > 0 {
		md.WriteHeading(2, "Aliases")
		md.WriteParagraph(wrapAndJoin("`", ", ", cmd.Aliases...))
	}

	// Write the long description
	if cmd.Long != "" {
		md.WriteHeading(2, "Description")
		md.WriteParagraph(cmd.Long)
	}

	// Write the usage
	if cmd.Runnable() {
		md.WriteHeading(2, "Usage")
		md.WriteCodeBlock(cmd.UseLine())
	}

	// Write the flags
	if cmd.Flags().HasFlags() {
		md.WriteHeading(2, "Flags")

		// Create a markdown table

		// Instantiate the headers
		headers := []string{"Flag", "Type", "Description", "Default"}

		// Instantiate the rows
		rows := [][]string{}
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			rows = append(rows, []string{
				"`--" + flag.Name + ", -" + flag.Shorthand + "`",
				wrap("`", flag.Value.Type()),
				flag.Usage,
				flag.DefValue,
			})
		})

		// Write the table
		md.WriteTable(headers, rows)
	}

	// Write the examples
	if len(cmd.Example) > 0 {
		md.WriteHeading(2, "Examples")
		md.WriteCodeBlock(cmd.Example)
	}

	// Write the see also section
	if cmd.HasParent() || len(cmd.Commands()) > 0 {
		md.WriteHeading(2, "See Also")
	}

	var name string = ""
	// Link to the parent command
	if cmd.HasParent() {
		name = cmd.Parent().Name()
		md.buf.WriteString("- ")
		md.WriteLink(name, "./"+name+".md")
		md.buf.WriteString("\n")

		// Link to the other sibling commands
		for _, subCmd := range cmd.Parent().Commands() {
			if subCmd != cmd {
				name = subCmd.Name()
				md.buf.WriteString("- ")
				md.WriteLink(name, "./"+name+".md")
				md.buf.WriteString("\n")
			}
		}
	}

	// Link to the child commands
	for _, subCmd := range cmd.Commands() {
		name = subCmd.Name()
		md.buf.WriteString("- ")
		md.WriteLink(name, "./"+name+".md")
		md.buf.WriteString("\n")
	}

	// Write the buffer to the writer
	_, err := md.buf.WriteTo(w)
	if err != nil {
		return err
	}

	return nil

}
