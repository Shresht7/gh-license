package api

// Reference: https://docs.github.com/en/rest/licenses?apiVersion=2022-11-28#get-all-commonly-used-licenses

// Fetch the list of all licenses from the GitHub API
func ListLicenses() ([]LicenseSimple, error) {
	endpoint := "licenses"
	return request[[]LicenseSimple](endpoint)
}
