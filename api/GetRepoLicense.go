package api

import (
	"github.com/cli/go-gh"
)

// Reference: https://docs.github.com/en/rest/licenses?apiVersion=2022-11-28#get-the-license-for-a-repository

// Fetch a license of a repository from the GitHub API
func GetRepoLicense(owner, repo string) (LicenseContent, error) {

	// Instantiate the gh client
	client, err := gh.RESTClient(nil)
	if err != nil {
		return LicenseContent{}, err
	}

	// Set the license endpoint
	endpoint := "repos/" + owner + "/" + repo + "/license"

	// Fetch the license from the GitHub API
	response := LicenseContent{} // Create a License struct to hold the response
	err = client.Get(endpoint, &response)
	if err != nil {
		return LicenseContent{}, err
	}

	// Return the license
	return response, nil

}
