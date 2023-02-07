package api

import (
	"github.com/cli/go-gh"
)

// Reference: https://docs.github.com/en/rest/licenses?apiVersion=2022-11-28#get-a-license

// Fetch a license from the GitHub API
func GetLicense(license string) (License, error) {

	// Instantiate the gh client
	client, err := gh.RESTClient(nil)
	if err != nil {
		return License{}, err
	}

	// Set the endpoint
	endpoint := "licenses/" + license

	// Fetch the license from the GitHub API
	response := License{} // Create a License struct to hold the response
	err = client.Get(endpoint, &response)
	if err != nil {
		return License{}, err
	}

	// Return the license information
	return response, nil

}
