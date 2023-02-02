package api

import "github.com/cli/go-gh"

// Reference: https://docs.github.com/en/rest/licenses?apiVersion=2022-11-28#get-all-commonly-used-licenses

// Fetch the list of all licenses from the GitHub API
func ListLicenses() ([]LicenseSimple, error) {

	// Instantiate the gh client
	client, err := gh.RESTClient(nil)
	if err != nil {
		return nil, err
	}

	// Set the endpoint
	endpoint := "licenses"

	// Fetch the list of licenses from the GitHub API
	response := []LicenseSimple{}
	err = client.Get(endpoint, &response)
	if err != nil {
		return nil, err
	}

	// Return the list of licenses
	return response, nil

}
