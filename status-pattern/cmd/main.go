package main

import (
	"github.com/echoloveyou/design-patterns/status-pattern"
)

func main() {
	gumballMachine := status_pattern.NewGumballMachine(2)
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
	gumballMachine.InsertQuarter()
	gumballMachine.TurnCrank()
}
