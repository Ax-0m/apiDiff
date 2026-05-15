package types

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
