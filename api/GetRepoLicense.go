package api

// Reference: https://docs.github.com/en/rest/licenses?apiVersion=2022-11-28#get-the-license-for-a-repository

// Fetch a license of a repository from the GitHub API
func GetRepoLicense(owner, repo string) (LicenseContent, error) {
	endpoint := "repos/" + owner + "/" + repo + "/license"
	return request[LicenseContent](endpoint)
}
