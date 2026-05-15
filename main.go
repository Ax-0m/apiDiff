package main

import (
	"fmt"
	"github.com/ax-0m/apidiff/parser"
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

	oldJSON, err := parser.ParseFile(oldFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	newJSON, err := parser.ParseFile(newFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Old JSON:", oldJSON)
	fmt.Println("New JSON:", newJSON)
}
