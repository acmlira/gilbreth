package structures

type Observer interface {
	Do(map[string]any) error
	GetID() string
}
