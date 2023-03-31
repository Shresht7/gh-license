package helpers

import (
	"bytes"
	"encoding/json"

	"github.com/cli/go-gh/pkg/jsonpretty"
)

// Prettify takes an input of any type and returns a prettified JSON string.
// If the input cannot be converted to JSON, or the JSON cannot be prettified, an error is returned.
func Prettify(input any) (string, error) {

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

	// Return the prettified JSON string
	return output.String(), nil

}
