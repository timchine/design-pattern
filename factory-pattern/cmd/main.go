package main

import (
	factory_pattern "github.com/echoloveyou/design-patterns/factory-pattern"
)

func main()  {
	factory_pattern.NewPizzaStore(new(factory_pattern.NYPizzaStory)).OrderPizza(factory_pattern.CheesePizzaType)
	factory_pattern.NewPizzaStore(new(factory_pattern.ChinaPizzaStory)).OrderPizza(factory_pattern.PepperoniPizzaType)
}
