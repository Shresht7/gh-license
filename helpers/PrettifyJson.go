package helpers

import (
	"bytes"
	"encoding/json"

	"github.com/cli/go-gh/pkg/jsonpretty"
)

// Prettifies JSON
func Prettify[T any](input T) (string, error) {
	// Convert the input to JSON
	json, err := json.Marshal(input)
	if err != nil {
		return "", err
	}

	// Prettify the JSON
	output := bytes.NewBuffer([]byte{})
	err = jsonpretty.Format(output, bytes.NewReader(json), "  ", true)
	if err != nil {
		return "", err
	}

	// Return the prettified JSON
	return output.String(), nil
}
