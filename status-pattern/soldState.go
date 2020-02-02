package status_pattern

import "fmt"

type SoldState struct {
	gumballMachine *GumballMachine
}

func NewSoldState(machine *GumballMachine) State {
	soldState := new(SoldState)
	soldState.gumballMachine = machine
	return soldState
}

func (s *SoldState) InsertQuarter() {
	fmt.Println("稍微等一下，你已经投入硬币了")
}

func (s *SoldState) EjectQuarter() {
	fmt.Println("你已经旋转了曲柄不能退回")
}

func (s *SoldState) TurnCrank() {
	fmt.Println("你转动了曲柄，但是不可以给你糖")
}

func (s *SoldState) Dispense() {
	s.gumballMachine.releaseBall()
	if s.gumballMachine.count > 0 {
		s.gumballMachine.SetState(s.gumballMachine.noQuarterState)
	} else {
		s.gumballMachine.SetState(s.gumballMachine.soldOutState)
	}
}
