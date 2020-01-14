package template_pattern

import "fmt"

type TaskTemplate struct {
	beforeTask func()
	afterTask func()
}

//定义执行任务的算法
func (task TaskTemplate) RunTask () {
	task.beforeTask()
	task.inTask()
	task.afterTask()
}

func (task TaskTemplate) inTask() {
	fmt.Println("在任务中")
}
