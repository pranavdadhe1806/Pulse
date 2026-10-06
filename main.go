package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Ensure at least one argument (the URL) is provided
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "error: no URL provided")
		os.Exit(1)
	}

	// Ensure no extra arguments are passed
	if len(os.Args) > 2 {
		fmt.Fprintln(os.Stderr, "error: More than one URL")
		os.Exit(1)
	}

	URL := os.Args[1]

	// Ensure the URL begins with "https://"
	if strings.HasPrefix(URL, "https://") {
		fmt.Println("URL: ", URL)
	} else {
		URL = "https://" + URL
		fmt.Println("URL: ", URL)
	}

}
