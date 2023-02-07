package api

// TODO: Implement Stringer interface for these structs

// Reference: https://docs.github.com/en/rest/licenses?apiVersion=2022-11-28

// Details of a particular license from the GitHub API
type License struct {
	Key            string   `json:"key"`            // The license key (e.g. "mit")
	Name           string   `json:"name"`           // The license name (e.g. "MIT License")
	Spdx_id        string   `json:"spdx_id"`        // The license SPDX ID (e.g. "MIT")
	Url            string   `json:"url"`            // The URL to the license on the web
	Node_id        string   `json:"node_id"`        // The node ID for the license
	Html_url       string   `json:"html_url"`       // The URL to the license on the web
	Description    string   `json:"description"`    // The description of the license
	Implementation string   `json:"implementation"` // The implementation of the license
	Permissions    []string `json:"permissions"`    // The permissions granted by the license
	Limitations    []string `json:"limitations"`    // The limitations of the license
	Conditions     []string `json:"conditions"`     // The conditions of the license
	Body           string   `json:"body"`           // The full text of the license
	Featured       bool     `json:"featured"`       // Whether or not the license is featured
}

// Simplified details of a particular license from the GitHub API
type LicenseSimple struct {
	Key     string `json:"key"`     // The license key (e.g. "mit")
	Name    string `json:"name"`    // The license name (e.g. "MIT License")
	Spdx_id string `json:"spdx_id"` // The license SPDX ID (e.g. "MIT")
	Url     string `json:"url"`     // The URL to the license on the web
	Node_id string `json:"node_id"` // The node ID for the license
}

// Details of a particular license from the GitHub API
type LicenseContent struct {
	Name         string `json:"name"`         // The license name (e.g. "MIT License")
	Path         string `json:"path"`         // The path to the license file
	Sha          string `json:"sha"`          // The SHA of the license file
	Size         int    `json:"size"`         // The size of the license file
	Url          string `json:"url"`          // The URL to the license file
	Html_url     string `json:"html_url"`     // The URL to the license file on the web
	Git_url      string `json:"git_url"`      // The URL to the license file in the Git repository
	Download_url string `json:"download_url"` // The URL to download the license file
	Type_        string `json:"type"`         // The type of the license file
	Content      string `json:"content"`      // The content of the license file
	Encoding     string `json:"encoding"`     // The encoding of the license file
	Links        struct {
		Self string `json:"self"` // The URL to the license file
		Git  string `json:"git"`  // The URL to the license file in the Git repository
		Html string `json:"html"` // The URL to the license file on the web
	} `json:"_links"` // The links to the license file
	License LicenseSimple `json:"license"` // The license details
}
