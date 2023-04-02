package api

// Reference: https://docs.github.com/en/rest/licenses?apiVersion=2022-11-28#get-the-license-for-a-repository

// Fetch a license of a repository from the GitHub API
func GetRepoLicense(owner, repo string) (LicenseContent, error) {

	// Set the endpoint
	endpoint := "repos/" + owner + "/" + repo + "/license"

	// Get the License information from the API
	response, err := request[LicenseContent](endpoint)
	if err != nil {
		return response, err
	}

	// Return the response
	return response, nil

}
