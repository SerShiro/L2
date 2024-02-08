package pattern

import "fmt"

type Element interface {
	Accept(visitor Visitor)
}

type ConcreteElementA struct {
	Name string
}

func (e *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(e)
}

type ConcreteElementB struct {
	Name string
}

func (e *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(e)
}

type Visitor interface {
	VisitConcreteElementA(element *ConcreteElementA)
	VisitConcreteElementB(element *ConcreteElementB)
}

type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Printf("ConcreteVisitor is visiting ConcreteElementA with name: %s\n", element.Name)
}

func (v *ConcreteVisitor) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Printf("ConcreteVisitor is visiting ConcreteElementB with name: %s\n", element.Name)
}

type ObjectStructure struct {
	elements []Element
}

func (os *ObjectStructure) Attach(element Element) {
	os.elements = append(os.elements, element)
}

func (os *ObjectStructure) Accept(visitor Visitor) {
	for _, element := range os.elements {
		element.Accept(visitor)
	}
}

func visitorStart() {
	elementA := &ConcreteElementA{Name: "ElementA"}
	elementB := &ConcreteElementB{Name: "ElementB"}

	visitor := &ConcreteVisitor{}

	objectStructure := &ObjectStructure{}
	objectStructure.Attach(elementA)
	objectStructure.Attach(elementB)

	objectStructure.Accept(visitor)
}
