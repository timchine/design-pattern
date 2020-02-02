package status_pattern

type State interface {
	InsertQuarter() //投入硬币
	EjectQuarter()  //退钱
	TurnCrank()     //旋转曲柄
	Dispense()      //发放糖果
}
