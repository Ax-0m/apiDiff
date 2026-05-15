package main

import (
	"fmt"
	"os"

	"github.com/ax-0m/apidiff/diff"
	"github.com/ax-0m/apidiff/parser"
	"github.com/ax-0m/apidiff/reporter"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: apidiff old.json new.json")
		os.Exit(1)
	}

	oldFile := os.Args[1]
	newFile := os.Args[2]

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

	changes := diff.Compare(oldJSON, newJSON, "")
	reporter.Report(changes)
}
