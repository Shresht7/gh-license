package helpers

import (
	"testing"
)

func TestListExamples(t *testing.T) {

	examples := []string{
		"gh license list",
		"gh license view mit",
		"gh license create mit",
		"gh license create mit --author \"John Doe\"",
		"gh license create mit --year 2020 --author \"John Doe\"",
	}

	result := ListExamples(examples)

	// ! Mind the 2 spaces before each example. This output is what looks good with cobra.
	expected := `gh license list
  gh license view mit
  gh license create mit
  gh license create mit --author "John Doe"
  gh license create mit --year 2020 --author "John Doe"
`

	if result != expected {
		t.Errorf("Expected \n %s \n\n got \n %s", expected, result)
	}
}
