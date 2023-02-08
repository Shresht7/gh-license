package main

import (
	"bytes"
	"io"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// GenerateMarkdown generates the markdown for the given command.
func GenerateMarkdown(cmd *cobra.Command, w io.Writer) error {

	// Instantiate a new buffer to store the markdown
	buf := new(bytes.Buffer)

	// Write the title
	buf.WriteString("# `" + cmd.Name() + "`\n\n")

	// Write the short description
	buf.WriteString(cmd.Short + "\n\n")

	// Write the aliases
	if len(cmd.Aliases) > 0 {
		buf.WriteString("## Aliases\n\n")
		buf.WriteString("`" + strings.Join(cmd.Aliases, "`, `") + "`\n\n")
	}

	// Write the long description
	if cmd.Long != "" {
		buf.WriteString("## Description\n\n")
		buf.WriteString(cmd.Long + "\n\n")
	}

	// Write the usage
	if cmd.Runnable() {
		buf.WriteString("## Usage\n\n")
		buf.WriteString("```sh\n")
		buf.WriteString(cmd.UseLine() + "\n")
		buf.WriteString("```\n\n")
	}

	// Write the flags
	if cmd.Flags().HasFlags() {
		buf.WriteString("## Flags\n\n")

		// Create a markdown table
		buf.WriteString("| Flag | Description | Default |\n")
		buf.WriteString("|------|-------------|---------|\n")

		// Write the flags to the table
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			buf.WriteString("| `--" + flag.Name + ", -" + flag.Shorthand + "` | " + flag.Usage + " | " + flag.DefValue + " |\n")
		})
		buf.WriteString("\n")
	}

	// Write the examples
	if len(cmd.Example) > 0 {
		buf.WriteString("## Examples\n\n")
		buf.WriteString("```sh\n")
		buf.WriteString(cmd.Example)
		buf.WriteString("```\n\n")
	}

	// Write the see also
	buf.WriteString("## See Also\n\n")

	// Link to the parent command
	if cmd.HasParent() {
		buf.WriteString("* [`" + cmd.Parent().Name() + "`](./" + cmd.Parent().Name() + ".md)\n")

		// Link to the other sibling commands
		for _, subCmd := range cmd.Parent().Commands() {
			if subCmd != cmd {
				buf.WriteString("* [`" + subCmd.Name() + "`](./" + subCmd.Name() + ".md)\n")
			}
		}
	}

	// Link to the child commands
	for _, subCmd := range cmd.Commands() {
		buf.WriteString("* [`" + subCmd.Name() + "`](./" + subCmd.Name() + ".md)\n")
	}

	// Write the buffer to the writer
	_, err := buf.WriteTo(w)
	if err != nil {
		return err
	}

	return nil

}
