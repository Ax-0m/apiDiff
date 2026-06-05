package cmd

import (
	"fmt"
	"github.com/Ax-0m/apiDiff/config"
)

func InitProject(projectName string) {
	err := config.CreateConfig(projectName)
	if err != nil {
		fmt.Print("error:", err)
		return
	}
}
