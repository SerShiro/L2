package pattern

import "fmt"

type ProductForFactory interface {
	Use()
}

type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() {
	fmt.Println("Using ConcreteProductA")
}

type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() {
	fmt.Println("Using ConcreteProductB")
}

type Creator interface {
	FactoryMethod() ProductForFactory
}

type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) FactoryMethod() ProductForFactory {
	return &ConcreteProductA{}
}

type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) FactoryMethod() ProductForFactory {
	return &ConcreteProductB{}
}

func Client(creator Creator) {
	product := creator.FactoryMethod()
	product.Use()
}

func factoryStart() {
	creatorA := &ConcreteCreatorA{}
	creatorB := &ConcreteCreatorB{}

	Client(creatorA)
	Client(creatorB)
}
