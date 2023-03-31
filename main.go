package main

import "github.com/Shresht7/gh-license/cmd"

// main is the entry point of the CLI application.
func main() {
	// Execute the root command of the CLI application.
	// This will parse the command line arguments and execute the appropriate command.
	// The root command is defined in cmd/root.go
	cmd.Execute()
}
