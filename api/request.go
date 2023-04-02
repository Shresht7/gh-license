package api

import (
	"github.com/cli/go-gh"
)

// request is a generic function to make a request to the GitHub API.
// It takes an endpoint and returns the response.
func request[T any](endpoint string) (T, error) {

	// Initialize the response struct with the correct type (T)
	response := new(T)

	// Get the gh client
	client, err := gh.RESTClient(nil)
	if err != nil {
		return *response, err
	}

	// Make the GET request
	err = client.Get(endpoint, &response)
	if err != nil {
		return *response, err
	}

	// Return the response
	return *response, nil

}
