package command_pattern

type OrderControl struct {
	orderQueue []Command
}

func (control *OrderControl) AddOrderCommand(command Command) {
	control.orderQueue = append(control.orderQueue, command)
}

func (control *OrderControl) Destroy() {
	for _, command := range control.orderQueue {
		command.Execute()
		control.orderQueue = control.orderQueue[1:]
	}
}