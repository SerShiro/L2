package pattern

import "fmt"

type Strategy interface {
	ExecuteStrategy()
}

type ConcreteStrategyA struct{}

func (s *ConcreteStrategyA) ExecuteStrategy() {
	fmt.Println("Executing ConcreteStrategyA")
}

type ConcreteStrategyB struct{}

func (s *ConcreteStrategyB) ExecuteStrategy() {
	fmt.Println("Executing ConcreteStrategyB")
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy() {
	if c.strategy != nil {
		c.strategy.ExecuteStrategy()
	} else {
		fmt.Println("No strategy is set")
	}
}

func strategyStart() {
	strategyA := &ConcreteStrategyA{}
	strategyB := &ConcreteStrategyB{}

	context := &Context{}

	context.SetStrategy(strategyA)
	context.ExecuteStrategy()

	context.SetStrategy(strategyB)
	context.ExecuteStrategy()
}
