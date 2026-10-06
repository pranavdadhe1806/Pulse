package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: URL Empty")
		return
	}

	if len(os.Args) > 2 {
		fmt.Println("Error: More than one URL")
		return
	}

	URL := os.Args[1]

	if strings.HasPrefix(URL, "https://") {
		fmt.Println("URL: ", URL)
	} else {
		URL := "https://" + URL
		fmt.Println("URL: ", URL)
	}
	// fmt.Println("URL: ", os.Args[1])

}
