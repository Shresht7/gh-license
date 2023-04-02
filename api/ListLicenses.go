package api

// Reference: https://docs.github.com/en/rest/licenses?apiVersion=2022-11-28#get-all-commonly-used-licenses

// Fetch the list of all licenses from the GitHub API
func ListLicenses() ([]LicenseSimple, error) {

	// Set the endpoint
	endpoint := "licenses"

	// Get the License information from the API
	response, err := request[[]LicenseSimple](endpoint)
	if err != nil {
		return response, err
	}

	// Return the response
	return response, nil

}
