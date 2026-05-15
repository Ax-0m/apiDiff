package reporter

import "fmt"

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
)

type Change struct {
	Path     string
	Type     string
	OldValue interface{}
	NewValue interface{}
}

func Report(changes []Change) {
	if len(changes) == 0 {
		fmt.Println("✔  no changes detected")
		return
	}

	for _, change := range changes {
		switch change.Type {
		case "added":
			fmt.Printf("%s✚  %s: added (%v)%s\n", colorGreen, change.Path, change.NewValue, colorReset)
		case "removed":
			fmt.Printf("%s✖  %s: removed (was %v)%s\n", colorRed, change.Path, change.OldValue, colorReset)
		case "modified":
			fmt.Printf("%s~  %s: changed  %v → %v%s\n", colorYellow, change.Path, change.OldValue, change.NewValue, colorReset)
		case "type_changed":
			fmt.Printf("%s⚠  %s: type changed  %T → %T%s\n", colorYellow, change.Path, change.OldValue, change.NewValue, colorReset)
		}
	}
}
