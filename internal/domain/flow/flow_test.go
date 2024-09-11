package flow_test

import (
	"testing"

	"github.com/acmlira/gilbreth/internal/domain/condition"
	"github.com/acmlira/gilbreth/internal/domain/flow"
	"github.com/acmlira/gilbreth/internal/domain/step"
)

func TestFlow(t *testing.T) {
	F := flow.New("my-flow")
	S1 := step.New("has key ", condition.IsPresent("my-key2"))
	S2 := step.New("has int")

	F.SetSteps(S1, S2)
	F.SetVariable("key", "my-key")
	F.SetVariable("key", "my-key2")
	F.SetVariable("key", "my-key2")
}
