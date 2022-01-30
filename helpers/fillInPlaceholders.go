package helpers

import (
	"strings"
)

//	Maps keys to their corresponding placeholder strings
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

//	Substitute placeholders with provided keys
func FillInPlaceholders(contents string, substitutions map[string]string) string {
	for key, placeholder := range placeholders {

		//	If no substitute, the continue iteration
		substitute, ok := substitutions[key]
		if !ok {
			continue
		}

		//	Replace placeholders
		for _, x := range placeholder {
			contents = strings.Replace(contents, x, substitute, -1)
		}
	}
	return contents
}
