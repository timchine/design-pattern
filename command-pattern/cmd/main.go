package main

import (
	command_pattern "github.com/echoloveyou/design-patterns/command-pattern"
	"time"
)

func main() {
	control := command_pattern.OrderControl{}
	for i := 0; i < 15; i++ {
		order := command_pattern.Order{
			GoodsId: "3",
			BuyNum:  1,
			OrderNo: "",
		}
		control.AddOrderCommand(command_pattern.QueueOrder{order})
	}
	go func() {
		control.Destroy()
	}()
	time.Sleep(10* time.Second)
}
