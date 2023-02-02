package helpers

// ListExamples takes a slice of examples and returns a raw string
func ListExamples(examples []string) string {
	var result string = examples[0] + "\n"
	for _, example := range examples[1:] {
		result += "  " + example + "\n"
	}
	return result
}
