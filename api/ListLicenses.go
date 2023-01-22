package api

import "github.com/cli/go-gh"

// Fetch the list of all licenses from the GitHub API
func ListLicenses() ([]LicenseName, error) {

	// Get gh client
	client, err := gh.RESTClient(nil)
	if err != nil {
		return nil, err
	}

	// License endpoint
	endpoint := "licenses"

	// Fetch the license
	response := []LicenseName{}
	err = client.Get(endpoint, &response)
	if err != nil {
		return nil, err
	}

	return response, nil

}
