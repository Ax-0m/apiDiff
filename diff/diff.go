package diff

import (
	"fmt"

	"github.com/Ax-0m/apiDiff/types"
)

func Compare(old, new map[string]interface{}, path string) []types.Change {
	var changes []types.Change

	for key, oldVal := range old {
		currentPath := buildPath(path, key)
		newVal, exists := new[key]

		if !exists {
			changes = append(changes, types.Change{
				Path:     currentPath,
				Type:     types.Removed,
				OldValue: oldVal,
			})
			continue
		}

		changes = append(changes, compareValues(currentPath, oldVal, newVal)...)
	}

	for key, newVal := range new {
		currentPath := buildPath(path, key)
		_, exists := old[key]

		if !exists {
			changes = append(changes, types.Change{
				Path:     currentPath,
				Type:     types.Added,
				NewValue: newVal,
			})
		}
	}

	return changes
}

func compareValues(path string, oldVal, newVal interface{}) []types.Change {
	var changes []types.Change

	oldType := fmt.Sprintf("%T", oldVal)
	newType := fmt.Sprintf("%T", newVal)

	if oldType != newType {
		return append(changes, types.Change{
			Path:     path,
			Type:     types.TypeChanged,
			OldValue: oldVal,
			NewValue: newVal,
		})
	}

	oldMap, oldIsMap := oldVal.(map[string]interface{})
	newMap, newIsMap := newVal.(map[string]interface{})

	if oldIsMap && newIsMap {
		return Compare(oldMap, newMap, path)
	}

	if oldVal != newVal {
		changes = append(changes, types.Change{
			Path:     path,
			Type:     types.Modified,
			OldValue: oldVal,
			NewValue: newVal,
		})
	}

	return changes
}

func buildPath(parent, key string) string {
	if parent == "" {
		return key
	}
	return parent + "." + key
}
