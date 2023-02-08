package main

import (
	"bytes"
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// GenerateMarkdown generates the markdown for the given command.
func GenerateMarkdown(cmd *cobra.Command, w io.Writer) error {

	// Instantiate a new buffer to store the markdown
	buf := new(bytes.Buffer)

	// Write the title
	buf.WriteString("# `" + cmd.CommandPath() + "`\n\n")

	// Write the short description
	buf.WriteString(cmd.Short + "\n\n")

	// Write the long description
	if cmd.Long != "" {
		buf.WriteString("## Description\n\n")
		buf.WriteString(cmd.Long + "\n\n")
	}

	// Write the aliases
	if len(cmd.Aliases) > 0 {
		buf.WriteString("## Aliases\n\n")
		buf.WriteString("```sh\n")
		buf.WriteString(cmd.Name() + " " + cmd.Aliases[0] + "\n")
		buf.WriteString("```\n\n")
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
		buf.WriteString("```sh\n")
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			buf.WriteString("--" + f.Name + " " + f.DefValue + " " + f.Usage + "\n")
		})
		buf.WriteString("```\n\n")
	}

	// Write the examples
	if len(cmd.Example) > 0 {
		buf.WriteString("## Examples\n\n")
		buf.WriteString("```sh\n")
		buf.WriteString(cmd.Example + "\n")
		buf.WriteString("```\n\n")
	}

	// Write the buffer to the writer
	_, err := buf.WriteTo(w)
	if err != nil {
		return err
	}

	return nil

}
