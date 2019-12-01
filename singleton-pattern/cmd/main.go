package main

import (
	"fmt"
	singleton_pattern "github.com/echoloveyou/design-patterns/singleton-pattern"
	"sync"
)

func main() {
	var (
		address  = "127.0.0.1:520"
		userName = "maYun"
		password = "123"
	)
	wg := sync.WaitGroup{}
	conn1 := singleton_pattern.ConnLazy(address, userName, password)
	conn2 := singleton_pattern.ConnLazy(address, userName, password)
	if conn1 == conn2 {
		fmt.Println("单线程单例没有问题")
	}
	wg.Add(1)
	go func() {
		conn3 := singleton_pattern.ConnLazy(address, userName, password)
		if conn3 == conn2 {
			fmt.Println("多协程没有问题")
		}
		wg.Done()
	}()
	wg.Wait()
}
