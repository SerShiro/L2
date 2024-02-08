package pattern

import "fmt"

type Handler interface {
	HandleRequest(request int)
	SetNext(handler Handler)
}

type ConcreteHandler struct {
	nextHandler Handler
	threshold   int
}

func NewConcreteHandler(threshold int) *ConcreteHandler {
	return &ConcreteHandler{threshold: threshold}
}

func (h *ConcreteHandler) SetNext(handler Handler) {
	h.nextHandler = handler
}

func (h *ConcreteHandler) HandleRequest(request int) {
	if request <= h.threshold {
		fmt.Printf("ConcreteHandler can handle the request %d\n", request)
	} else if h.nextHandler != nil {
		fmt.Printf("ConcreteHandler passes the request %d to the next handler\n", request)
		h.nextHandler.HandleRequest(request)
	} else {
		fmt.Printf("ConcreteHandler cannot handle the request %d\n", request)
	}
}

func chainStart() {
	handler1 := NewConcreteHandler(10)
	handler2 := NewConcreteHandler(20)
	handler3 := NewConcreteHandler(30)

	handler1.SetNext(handler2)
	handler2.SetNext(handler3)

	handler1.HandleRequest(5)
	handler1.HandleRequest(25)
	handler1.HandleRequest(35)
}
