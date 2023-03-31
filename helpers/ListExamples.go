package helpers

// ListExamples takes a slice of examples (strings) and returns
// a string containing the examples, each on a new line, indented by two spaces.
// This is used to display examples in the help text for commands.
func ListExamples(examples []string) string {
	var result string
	for _, example := range examples {
		result += "  " + example + "\n"
	}
	return result
}
