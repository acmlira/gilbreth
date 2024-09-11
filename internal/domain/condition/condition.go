package condition

type condition struct {
	key     string
	X, Y, Z any
}

type Condition interface {
	Check(variables map[string]any) bool
}
