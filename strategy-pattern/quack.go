package strategy_pattern

import "fmt"

type QuackBehavior interface {
	Quack()
}

type Quack struct {

}

type MuteQuack struct {

}

func (Quack) Quack() {
	fmt.Println("我是会叫")
}

func (MuteQuack) Quack() {
	fmt.Println("我是一个哑的鸭")
}


