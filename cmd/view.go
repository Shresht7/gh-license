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

		//	TODO: Display information selectively based on passed arguments. e.g. gh license view mit description implementation

	},
}

func init() {
	rootCmd.AddCommand(viewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// viewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// viewCmd.Flags().BoolVarP(&web, "web", "w", false, "View in browser")
}
