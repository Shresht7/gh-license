package cmd

import (
	"fmt"

	"github.com/cli/go-gh"
	"github.com/spf13/cobra"
)

type LicenseInfo struct {
	Name           string
	Spdx_id        string
	Html_url       string
	Description    string
	Implementation string
	Permissions    []string
	Limitations    []string
	Conditions     []string
	Body           string
}

//	============
//	VIEW COMMAND
//	============

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View details about a particular license",
	Long:  `View details about a particular license`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var licenseID string = args[0]

		//	Get gh client
		client, err := gh.RESTClient(nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		//	License endpoint
		license := "licenses/" + licenseID

		//	Fetch the license
		response := LicenseInfo{}
		err = client.Get(license, &response)
		if err != nil {
			fmt.Println("Invalid License")
			return
		}

		//	Print information about the License
		fmt.Println("\nName:", response.Name)
		fmt.Println("SPDX_ID:", response.Spdx_id)
		fmt.Println("URL:", response.Html_url)
		fmt.Println("\nDescription:\n\n", response.Description)
		fmt.Println("\nImplementation:\n\n", response.Implementation)
		fmt.Println("\nPermissions:", response.Permissions)
		fmt.Println("Conditions:", response.Conditions)
		fmt.Println("Limitations:", response.Limitations)
		fmt.Println("\nBody:\n\n", response.Body)

	},
}

func init() {
	//	Add view command
	rootCmd.AddCommand(viewCmd)
}
