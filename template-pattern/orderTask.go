package template_pattern

import "fmt"

type OrderTask struct {
	TaskTemplate
}

func NewOrderTask() *OrderTask {
	var task = new(OrderTask)
	task.beforeTask = func() {
		fmt.Println("查找订单")
	}
	task.afterTask = func() {
		fmt.Println("执行结束")
	}
	return task
}