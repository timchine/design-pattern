package status_pattern

import "fmt"

type HasQuarterState struct {
	gumballMachine *GumballMachine
}

func NewHasQuarterState(machine *GumballMachine) State {
	hasQuarterState := new(HasQuarterState)
	hasQuarterState.gumballMachine = machine
	return hasQuarterState
}

func (h *HasQuarterState) InsertQuarter() {
	fmt.Println("你不能再次投入硬币")
}

func (h *HasQuarterState) EjectQuarter() {
	fmt.Println("你的硬币已被退回")
	h.gumballMachine.SetState(h.gumballMachine.noQuarterState)
}

func (h *HasQuarterState) TurnCrank() {
	fmt.Println("你旋转了曲柄")
	h.gumballMachine.SetState(h.gumballMachine.soldState)
}

func (h *HasQuarterState) Dispense() {
	fmt.Println("不能释放糖果")
}