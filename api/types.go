package api

// Details about a particular license from the GitHub API
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

// Name of a license from the GitHub API
type LicenseName struct {
	Key     string `json:"key"`     // The license key (e.g. "mit")
	Name    string `json:"name"`    // The license name (e.g. "MIT License")
	Spdx_id string `json:"spdx_id"` // The license SPDX ID (e.g. "MIT")
	Url     string `json:"url"`     // The URL to the license on the web
	Node_id string `json:"node_id"` // The node ID for the license
}
