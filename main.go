package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		//	TODO: Show Help Message
		fmt.Println("Please provide an argument")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "list":
		list()
	}
}
