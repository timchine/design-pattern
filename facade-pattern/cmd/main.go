package main

import facade_pattern "github.com/echoloveyou/design-patterns/facade-pattern"

func main()  {
	var orderService facade_pattern.OrderService
	orderService.SubOrder()
}
