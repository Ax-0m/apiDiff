package cmd

import (
	"fmt"
	"github.com/Ax-0m/apiDiff/config"
)

func InitProject(projectName, baseURL string, endpoints []string) {
	err := config.CreateConfig(projectName, baseURL, endpoints)
	if err != nil {
		fmt.Print("error:", err)
		return
	}
}
