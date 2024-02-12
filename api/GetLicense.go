package api

// Reference: https://docs.github.com/en/rest/licenses?apiVersion=2022-11-28#get-a-license

// Fetch a license from the GitHub API
func GetLicense(license string) (License, error) {
	endpoint := "licenses/" + license
	return request[License](endpoint)
}
