package api

// Reference: https://docs.github.com/en/rest/licenses?apiVersion=2022-11-28#get-a-license

// Fetch a license from the GitHub API
func GetLicense(license string) (License, error) {

	// Set the endpoint for the request
	endpoint := "licenses/" + license

	// Get the License information from the API
	response, err := request[License](endpoint)
	if err != nil {
		return response, err
	}

	// Return the response
	return response, nil

}
