package reporter

import (
	"fmt"

	"github.com/Ax-0m/apiDiff/types"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
)

func Report(changes []types.Change) {
	if len(changes) == 0 {
		fmt.Println("✔  no changes detected")
		return
	}

	for _, change := range changes {
		switch change.Type {
		case types.Added:
			fmt.Printf("%s✚  %s: added (%v)%s\n", colorGreen, change.Path, change.NewValue, colorReset)
		case types.Removed:
			fmt.Printf("%s✖  %s: removed (was %v)%s\n", colorRed, change.Path, change.OldValue, colorReset)
		case types.Modified:
			fmt.Printf("%s~  %s: changed  %v → %v%s\n", colorYellow, change.Path, change.OldValue, change.NewValue, colorReset)
		case types.TypeChanged:
			fmt.Printf("%s⚠  %s: type changed  %s → %s%s\n", colorYellow, change.Path, friendlyType(change.OldValue), friendlyType(change.NewValue), colorReset)
		}
	}
}

func friendlyType(v interface{}) string {
	switch v.(type) {
	case float64:
		return "number"
	case string:
		return "string"
	case bool:
		return "boolean"
	case map[string]interface{}:
		return "object"
	case []interface{}:
		return "array"
	case nil:
		return "null"
	default:
		return "unknown"
	}
}
