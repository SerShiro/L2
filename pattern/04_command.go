package pattern

import "fmt"

type Command interface {
	Execute()
}

type Receiver struct {
	Name string
}

func (r *Receiver) Action() {
	fmt.Printf("Receiver %s is performing the action\n", r.Name)
}

type ConcreteCommand struct {
	receiver *Receiver
}

func NewConcreteCommand(receiver *Receiver) *ConcreteCommand {
	return &ConcreteCommand{receiver: receiver}
}

func (cc *ConcreteCommand) Execute() {
	cc.receiver.Action()
}

type Invoker struct {
	command Command
}

func (i *Invoker) SetCommand(command Command) {
	i.command = command
}

func (i *Invoker) ExecuteCommand() {
	if i.command != nil {
		i.command.Execute()
	} else {
		fmt.Println("No command is set")
	}
}

func commandStart() {
	receiver := &Receiver{Name: "ReceiverA"}

	commandA := NewConcreteCommand(receiver)
	commandB := NewConcreteCommand(receiver)

	invoker := &Invoker{}

	invoker.SetCommand(commandA)

	invoker.ExecuteCommand()

	invoker.SetCommand(commandB)
	invoker.ExecuteCommand()
}
