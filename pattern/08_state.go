package pattern

import "fmt"

type ContextForState struct {
	state State
}

type State interface {
	Handle(context *ContextForState)
}

type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle(context *ContextForState) {
	fmt.Println("Handling state A")
	context.state = &ConcreteStateB{}
}

type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle(context *ContextForState) {
	fmt.Println("Handling state B")
	context.state = &ConcreteStateA{}
}

func stateStart() {
	context := &ContextForState{state: &ConcreteStateA{}}

	context.state.Handle(context)

	context.state.Handle(context)
}
