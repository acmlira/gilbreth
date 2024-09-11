package step

import (
	"errors"

	"github.com/acmlira/gilbreth/internal/domain/condition"
	"github.com/google/uuid"
)

type Step interface {
	Do(map[string]any) error
	GetID() string
}

type step struct {
	ID         uuid.UUID
	Name       string
	conditions []condition.Condition
}

func (s *step) Do(variables map[string]any) error {
	for _, cd := range s.conditions {
		if c := cd.Check(variables); !c {
			return errors.New("step does not match the condition")
		}
	}

	return nil
}

func (s *step) GetID() string {
	return s.ID.String()
}

func New(name string, conditions ...condition.Condition) Step {
	return &step{
		ID:         uuid.New(),
		Name:       name,
		conditions: conditions,
	}
}
