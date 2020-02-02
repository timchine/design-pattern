package status_pattern

import "fmt"

type GumballMachine struct {
	noQuarterState  State //无硬币
	hasQuarterState State //有硬币
	soldState       State //售出
	soldOutState    State //售空
	count           int
	state           State //当前执行操作对应的状态
}

//new一个GumballMachine的对象
func NewGumballMachine(count int) *GumballMachine {
	gumballMachine := new(GumballMachine)
	//设置糖果的数量以免超卖
	gumballMachine.count = count
	//new一个NoQuarterState
	gumballMachine.noQuarterState = NewNoQuarterState(gumballMachine)
	gumballMachine.hasQuarterState = NewHasQuarterState(gumballMachine)
	gumballMachine.soldState = NewSoldState(gumballMachine)
	gumballMachine.soldOutState = NewSoldOutState(gumballMachine)
	//如果糖果数量大于1则将状态设为没有硬币状态
	if gumballMachine.count > 0 {
		gumballMachine.state = gumballMachine.noQuarterState
	}
	return gumballMachine
}

func (m *GumballMachine) releaseBall() {
	fmt.Println("释放糖果")
	if m.count != 0 {
		m.count = m.count - 1
	}
}

func (m *GumballMachine) InsertQuarter() {
	m.state.InsertQuarter()
}

func (m *GumballMachine) EjectQuarter() {
	m.state.EjectQuarter()
}

func (m *GumballMachine) TurnCrank() {
	m.state.TurnCrank()
	m.state.Dispense()
}

func (m *GumballMachine) SetState(state State) {
	m.state = state
}
