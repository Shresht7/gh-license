package cmd

import (
	"fmt"

	"github.com/cli/go-gh"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a license file",
	Long:  `Create a license file for your project`,
	Run: func(cmd *cobra.Command, args []string) {

		//	Get gh client
		client, err := gh.RESTClient(nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		license := "licenses/" + args[0]

		//	Fetch the license
		response := struct{ Body string }{}
		err = client.Get(license, &response)
		if err != nil {
			fmt.Println(err)
			return
		}

		//	Print the list of licenses
		fmt.Println(response.Body)

	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
