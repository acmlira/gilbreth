package flow

import (
	"github.com/acmlira/gilbreth/internal/domain/step"
	"github.com/acmlira/gilbreth/internal/structures"
	"github.com/google/uuid"
)

type flow struct {
	id        uuid.UUID
	steps     structures.Queue
	variables map[string]any
	key       string
	observers []structures.Observer
}

type Flow interface {
	GetID() uuid.UUID
	GetVariables() map[string]any
	GetVariable(string) any
	SetVariable(key string, value any)
	SetSteps(steps ...step.Step)
}

func (f *flow) GetID() uuid.UUID {
	return f.id
}

func (f *flow) GetVariables() map[string]any {
	return f.variables
}

func (f *flow) GetVariable(key string) any {
	return f.variables[key]
}

func (f *flow) SetVariable(key string, value any) {
	f.variables[key] = value
	f.notifyFirst()
}

var flowsIndexedByID = make(map[uuid.UUID]Flow)
var flowsIndexedByKey = make(map[string]Flow)

func New(key string) Flow {
	f := new(flow)
	f.id = uuid.New()
	f.variables = make(map[string]any)
	f.steps = structures.NewQueue()
	f.key = key

	// Save, maybe move to a repository (?)
	flowsIndexedByID[f.id] = f
	flowsIndexedByKey[key] = f

	return f
}

func GetByID(id uuid.UUID) Flow {
	return flowsIndexedByID[id]
}

func GetByKey(key string) Flow {
	return flowsIndexedByKey[key]
}

func (f *flow) register(observer structures.Observer) {
	f.observers = append(f.observers, observer)
}

func (f *flow) deregister(observerToRemove structures.Observer) {
	observerListLength := len(f.observers)
	for i, observer := range f.observers {
		if observerToRemove.GetID() == observer.GetID() {
			f.observers[observerListLength-1], f.observers[i] = f.observers[i], f.observers[observerListLength-1]
			f.observers = f.observers[:observerListLength-1]
			return
		}
	}
}

func (f *flow) notifyFirst() {
	if !f.steps.IsEmpty() {
		if s, ok := f.steps.Head().(step.Step); ok {
			if s.Do(f.variables) == nil {
				f.deregister(s)
				f.steps.Dequeue()
			}
		}
	}
}

func (f *flow) SetSteps(steps ...step.Step) {
	for _, s := range steps {
		f.steps.Enqueue(s)
		f.register(s)
	}
}
