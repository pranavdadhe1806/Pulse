package main

import (
	"fmt"
	"os"
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

	fmt.Println("URL: ", os.Args[1])

	//URL:= os.Args[1]

}
