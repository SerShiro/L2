package pattern

import "fmt"

type SubsystemA struct {
}

func (s *SubsystemA) OperationA() {
	fmt.Println("SubsystemA: OperationA")
}

type SubsystemB struct {
}

func (s *SubsystemB) OperationB() {
	fmt.Println("SubsystemB: OperationB")
}

type SubsystemC struct {
}

func (s *SubsystemC) OperationC() {
	fmt.Println("SubsystemC: OperationC")
}

type SubsystemD struct {
}

func (s *SubsystemD) OperationD() {
	fmt.Println("SubsystemD: OperationD")
}

type Facade struct {
	subsystemA *SubsystemA
	subsystemB *SubsystemB
	subsystemC *SubsystemC
	subsystemD *SubsystemD
}

func NewFacade() *Facade {
	return &Facade{
		subsystemA: &SubsystemA{},
		subsystemB: &SubsystemB{},
		subsystemC: &SubsystemC{},
		subsystemD: &SubsystemD{},
	}
}

func (f *Facade) OperationFacade() {
	fmt.Println("Начало работы...")
	f.subsystemA.OperationA()
	f.subsystemB.OperationB()
	f.subsystemC.OperationC()
	f.subsystemD.OperationD()
	fmt.Println("Конец работы!")
}

func facadeStart() {
	facade := NewFacade()
	facade.OperationFacade()
}
