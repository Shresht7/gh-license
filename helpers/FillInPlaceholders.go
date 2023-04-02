package helpers

import (
	"strings"
)

// Maps keys to placeholder strings. The placeholders are replaced with the values of the corresponding key.
var placeholders = map[string][]string{
	"author": {
		"[fullname]",
		"{fullname}",
		"<name of author>",
		"{name of author}",
		"{name of copyright owner}",
	},
	"project": {
		"[project]",
		"{project}",
	},
	"description": {
		"{one line to give the program's name and a brief idea of what it does.}",
		"<one line to give the program's name and a brief idea of what it does.>",
	},
	"year": {
		"[year]",
		"<year>",
		"{year}",
		"[yyyy]",
		"{yyyy}",
	},
}

// FillInPlaceholders replaces all placeholders in the provided contents with the provided substitutions.
func FillInPlaceholders(contents string, substitutions map[string]string) string {

	// Iterate over placeholders and substitute them with the provided substitutions
	for key, placeholder := range placeholders {

		// Get the substitution for the current placeholder
		substitute, ok := substitutions[key]
		// If no substitution is provided, skip this placeholder
		if !ok {
			continue
		}

		// Replace all placeholders with the substitution
		for _, x := range placeholder {
			contents = strings.Replace(contents, x, substitute, -1)
		}

	}

	// Return the contents with the placeholders replaced
	return contents

}
