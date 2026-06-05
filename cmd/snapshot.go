package cmd

import (
	"fmt"
	"github.com/Ax-0m/apiDiff/config"
	"github.com/Ax-0m/apiDiff/fetcher"
	"github.com/Ax-0m/apiDiff/snapshot"
)

func SnapshotProject(projectName string) {
	fmt.Println("debug: SnapshotProject called with", projectName)
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("error loading config:", err)
		return
	}

	for _, endpoint := range cfg.Endpoints {
		url := cfg.BaseURL + endpoint
		fmt.Println("fetching:", url)

		data, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println("error fetching endpoint:", endpoint, err)
			continue
		}

		err = snapshot.Save(projectName, endpoint, data)
		if err != nil {
			fmt.Println("error saving snapshot:", endpoint, err)
			continue
		}
	}

	fmt.Println("snapshots saved for project:", projectName)
}
