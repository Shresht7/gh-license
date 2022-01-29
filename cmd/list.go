package cmd

import (
	"fmt"

	"github.com/cli/go-gh"
	"github.com/spf13/cobra"
)

type License struct {
	Key     string
	Spdx_id string
	Name    string
	Url     string
	Node_id string
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List licenses",
	Long:  `Show a list of licenses`,
	Run: func(cmd *cobra.Command, args []string) {

		//	Get gh client
		client, err := gh.RESTClient(nil)
		if err != nil {
			fmt.Println(err)
			return
		}

		//	Fetch the list of licenses
		response := []License{}
		err = client.Get("licenses", &response)
		if err != nil {
			fmt.Println(err)
			return
		}

		//	Print the list of licenses
		for _, license := range response {
			fmt.Printf("%-16s %s\n", license.Spdx_id, license.Name)
		}

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
