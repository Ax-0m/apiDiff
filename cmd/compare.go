package cmd

import (
	"fmt"
	"github.com/Ax-0m/apiDiff/config"
	"github.com/Ax-0m/apiDiff/diff"
	"github.com/Ax-0m/apiDiff/fetcher"
	"github.com/Ax-0m/apiDiff/reporter"
	"github.com/Ax-0m/apiDiff/snapshot"
)

func CompareProject(projectname string) {
	cfg, err := config.LoadConfig()

	if err != nil {
		fmt.Println("error loading config: ", err)
		return
	}

	for _, endpoint := range cfg.Endpoints {
		url := cfg.BaseURL + endpoint
		fmt.Println("comparing: ", url)

		freshdata, err := fetcher.Fetch(url)

		if err != nil {
			fmt.Println("error fetching endpoint: ", endpoint, err)
			continue
		}

		savedData, err := snapshot.Load(projectname, endpoint)

		if err != nil {
			fmt.Println("error loading snapshot: ", endpoint, err)
			continue
		}

		fmt.Println("\n--- diff for", endpoint, "---")
		changes := diff.Compare(savedData, freshdata, "")
		reporter.Report(changes)

	}
}
