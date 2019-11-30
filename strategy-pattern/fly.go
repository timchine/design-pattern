package strategy_pattern

import "fmt"

//定义飞行行为的接口
type FlyBehavior interface {
	Fly()
}

type FlyWithWings struct {

}

type FlyNoWay struct {

}

func (fly FlyNoWay) Fly() {
	fmt.Println("我不会飞")
}

func (fly FlyWithWings) Fly() {
	fmt.Println("我是一个会飞的鸭子")
}