package helpers

import (
	"fmt"
	"testing"
)

var substitutions = map[string]string{
	"author":      "John Doe",
	"project":     "My Project",
	"description": "This is a description",
	"year":        "2019",
}

func TestFillInPlaceholders(t *testing.T) {

	// Setup test cases and expected results
	testCases := []struct {
		contents    string
		expected    string
		description string
	}{
		{
			contents:    "My name is [fullname]",
			expected:    fmt.Sprintf("My name is %s", substitutions["author"]),
			description: "Replace [fullname] placeholder",
		},
		{
			contents:    "My name is {fullname}",
			expected:    fmt.Sprintf("My name is %s", substitutions["author"]),
			description: "Replace {fullname} placeholder",
		},
		{
			contents:    "My name is <name of author>",
			expected:    fmt.Sprintf("My name is %s", substitutions["author"]),
			description: "Replace <name of author> placeholder",
		},
		{
			contents:    "My name is {name of author}",
			expected:    fmt.Sprintf("My name is %s", substitutions["author"]),
			description: "Replace {name of author} placeholder",
		},
		{
			contents:    "My name is {name of copyright owner}",
			expected:    fmt.Sprintf("My name is %s", substitutions["author"]),
			description: "Replace {name of copyright owner} placeholder",
		},
		{
			contents:    "My project is [project]",
			expected:    fmt.Sprintf("My project is %s", substitutions["project"]),
			description: "Replace [project] placeholder",
		},
		{
			contents:    "My project is {project}",
			expected:    fmt.Sprintf("My project is %s", substitutions["project"]),
			description: "Replace {project} placeholder",
		},
		{
			contents:    "This is a description",
			expected:    fmt.Sprintf("This is a description"),
			description: "No placeholders",
		},
		{
			contents:    "This is a description {one line to give the program's name and a brief idea of what it does.}",
			expected:    fmt.Sprintf("This is a description %s", substitutions["description"]),
			description: "Replace {one line to give the program's name and a brief idea of what it does.} placeholder",
		},
		{
			contents:    "This is a description <one line to give the program's name and a brief idea of what it does.>",
			expected:    fmt.Sprintf("This is a description %s", substitutions["description"]),
			description: "Replace <one line to give the program's name and a brief idea of what it does.> placeholder",
		},
		{
			contents:    "This year is [year]",
			expected:    fmt.Sprintf("This year is %s", substitutions["year"]),
			description: "Replace [year] placeholder",
		},
		{
			contents:    "This year is {year}",
			expected:    fmt.Sprintf("This year is %s", substitutions["year"]),
			description: "Replace {year} placeholder",
		},
		{
			contents:    "This year is <year>",
			expected:    fmt.Sprintf("This year is %s", substitutions["year"]),
			description: "Replace <year> placeholder",
		},
	}

	// Run test cases
	for _, testCase := range testCases {
		actual := FillInPlaceholders(testCase.contents, substitutions)
		if actual != testCase.expected {
			t.Errorf("Test failed: %s - Expected: %s, got: %s", testCase.description, testCase.expected, actual)
		}
	}

}
