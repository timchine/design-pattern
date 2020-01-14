package main

import template_pattern "github.com/echoloveyou/design-patterns/template-pattern"

func main() {
	template_pattern.NewCarTask().RunTask()
	template_pattern.NewOrderTask().RunTask()
}