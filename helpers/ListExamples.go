package helpers

// ListExamples takes a slice of examples and returns a raw string
func ListExamples(examples []string) string {
	var result string
	for _, example := range examples {
		result += "  " + example + "\n"
	}
	return result
}
