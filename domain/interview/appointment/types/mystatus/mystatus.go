package mystatus

import "fmt"

const (
	Todo       = "To Do"
	InProgress = "In Progress"
	Done       = "Done"
)

type Type struct {
	Value string
}

func NewType(value string) (Type, error) {
	if value != Todo && value != InProgress && value != Done {
		return Type{}, fmt.Errorf("invalid status: %s", value)
	}
	return Type{Value: value}, nil
}
