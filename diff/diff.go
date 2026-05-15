package diff

import "fmt"

type ChangeType string

const (
	Added       ChangeType = "added"
	Removed     ChangeType = "removed"
	Modified    ChangeType = "modified"
	TypeChanged ChangeType = "type_changed"
)

type Change struct {
	Path     string
	Type     ChangeType
	OldValue interface{}
	NewValue interface{}
}

func Compare(old, new map[string]interface{}, path string) []Change {
	var changes []Change

	for key, oldVal := range old {
		currentPath := buildPath(path, key)
		newVal, exists := new[key]

		if !exists {
			changes = append(changes, Change{
				Path:     currentPath,
				Type:     Removed,
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
			changes = append(changes, Change{
				Path:     currentPath,
				Type:     Added,
				NewValue: newVal,
			})
		}
	}
	return changes
}

func compareValues(path string, oldVal, newVal interface{}) []Change {
	var changes []Change

	oldType := fmt.Sprintf("%T", oldVal)
	newType := fmt.Sprintf("%T", newVal)

	if oldType != newType {
		return append(changes, Change{
			Path:     path,
			Type:     TypeChanged,
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
		changes = append(changes, Change{
			Path:     path,
			Type:     Modified,
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
