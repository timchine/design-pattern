package status_pattern

import "fmt"

type NoQuarterState struct {
	gumballMachine *GumballMachine
}

func NewNoQuarterState(machine *GumballMachine) State {
	noQuarterState := new(NoQuarterState)
	noQuarterState.gumballMachine = machine
	return noQuarterState
}

//投入硬币后将状态改为有硬币这样就看似修改了类。
func (n *NoQuarterState) InsertQuarter() {
	fmt.Println("你投入了硬币")
	n.gumballMachine.SetState(n.gumballMachine.hasQuarterState)
}

func (n *NoQuarterState) EjectQuarter() {
	fmt.Println("你未投入硬币")
}

func (n *NoQuarterState) TurnCrank() {
	fmt.Println("你转动了曲柄，但是没有硬币")
}

func (n *NoQuarterState) Dispense() {
	fmt.Println("你需要先投入硬币")
}
