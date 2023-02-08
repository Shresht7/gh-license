package main

import (
	"log"
	"os"

	"github.com/spf13/cobra/doc"

	"github.com/Shresht7/gh-license/cmd"
)

// Generate generates the documentation for the given package.
func main() {

	// for each command in cmd create a markdown file in docs
	for _, cmd := range cmd.RootCmd.Commands() {

		// create a file in docs
		w, err := os.Create("docs/gh-license-" + cmd.Name() + ".md")
		if err != nil {
			log.Fatal(err)
		}
		defer w.Close()

		// generate markdown for the command
		err = doc.GenMarkdown(cmd, w)
		if err != nil {
			log.Fatal(err)
		}

	}
}
