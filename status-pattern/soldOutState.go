package status_pattern

import "fmt"

type SoldOutState struct {
	gumballMachine *GumballMachine
}

func NewSoldOutState(machine *GumballMachine) State {
	soldOutState := new(SoldOutState)
	soldOutState.gumballMachine = machine
	return soldOutState
}

func (s *SoldOutState) InsertQuarter() {
	fmt.Println("糖果售空")
}

func (s *SoldOutState) EjectQuarter() {
	fmt.Println("不可以退回硬币")
}

func (s *SoldOutState) TurnCrank() {
	fmt.Println("你旋转了曲柄但是糖果售空")
}

func (s *SoldOutState) Dispense() {
	fmt.Println("没有糖果释放")
}

