package template_pattern

import "fmt"

type CarTask struct {
	TaskTemplate
}

func NewCarTask() *CarTask {
	var task = new(CarTask)
	task.beforeTask = func() {
			fmt.Println("查找购物车")
		}
	task.afterTask = func() {
		fmt.Println("执行结束")
	}
	return task
}