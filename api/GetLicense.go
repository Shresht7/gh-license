package api

import (
	"github.com/cli/go-gh"
)

// Fetch a license from the GitHub API
func GetLicense(license string) (License, error) {

	// Get gh client
	client, err := gh.RESTClient(nil)
	if err != nil {
		return License{}, err
	}

	// License endpoint
	endpoint := "licenses/" + license

	// Fetch the license
	response := License{}
	err = client.Get(endpoint, &response)
	if err != nil {
		return License{}, err
	}

	return response, nil

}
