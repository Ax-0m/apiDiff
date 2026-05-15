package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 { // os.args is an array
		fmt.Println("Usage: apidiff old.json new.json")
		os.Exit(1)
	}
	oldFile := os.Args[1]
	newFile := os.Args[2]

	fmt.Println("Old:", oldFile)
	fmt.Println("New:", newFile)
}
