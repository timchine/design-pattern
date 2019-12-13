package adapter_pattern

import "fmt"

type Adaptee struct {

}

func (adaptee Adaptee) EspecialRequest() {
	fmt.Println("我是一个特殊的请求")
	fmt.Println("一堆biubiu逻辑")
}

//上面定义了被装饰者。可以看到有一堆的逻辑
type Adapter interface {
	Request()
}

type Adapt struct {
	adaptee *Adaptee
}

func (adapt *Adapt)SetAdaptee(adaptee *Adaptee)  {
	adapt.adaptee = adaptee
}

func (a *Adapt) Request() {
	a.adaptee.EspecialRequest()
}

