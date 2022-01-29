package main

import (
	"fmt"

	"github.com/cli/go-gh"
)

type License struct {
	Key     string
	Spdx_id string
	Name    string
	Url     string
	Node_id string
}

func list() {

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

}
