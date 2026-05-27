package main

import (
	"fmt"
	"os"

	"github.com/Ax-0m/apiDiff/cmd"
	"github.com/Ax-0m/apiDiff/diff"
	"github.com/Ax-0m/apiDiff/parser"
	"github.com/Ax-0m/apiDiff/reporter"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: apidiff <command> [args]")
		fmt.Println("Commands: init, snapshot, compare")
		fmt.Println("Or: apidiff old.json new.json")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "init":
		if len(os.Args) < 3 {
			fmt.Println("Usage: apidiff init <project-name>")
			os.Exit(1)
		}
		cmd.InitProject(os.Args[2])
	default:
		if len(os.Args) < 3 {
			fmt.Println("Usage: apidiff old.json new.json")
			os.Exit(1)
		}
		oldJSON, err := parser.ParseFile(os.Args[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		newJSON, err := parser.ParseFile(os.Args[2])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		changes := diff.Compare(oldJSON, newJSON, "")
		reporter.Report(changes)
	}
}
