package factory_pattern

import "fmt"

type SimplePizzaFactory interface {
	NewSimplePizza(SimplePizzaType) SimplePizza
}

type SimplePizzaType int8

const (
	_                    SimplePizzaType = iota
	cheeseSimplePizza                    //奶酪披萨
	pepperoniSimplePizza                 //香肠披萨
)

type ISimplePizza interface {
	Box() //
}

//定义一个简单披萨的超类
type SimplePizza struct {
	Name  string
	Price float64
}

func NewSimplePizza(t SimplePizzaType) ISimplePizza {
	pizza := SimplePizza{}
	return pizza.NewSimplePizza(t)
}

//根据传入类型不同返回不同的披萨
func (s SimplePizza) NewSimplePizza(t SimplePizzaType) ISimplePizza {
	switch t {
	case cheeseSimplePizza:
		return new(CheeseSimplePizza)
	case pepperoniSimplePizza:
		return new(PepperoniSimplePizza)
	default:
		return new(SimplePizza)
	}
}

func (s SimplePizza) Box() {
	fmt.Println("披萨打包好了")
}

//奶酪披萨
type CheeseSimplePizza struct {
	SimplePizza
}

//香肠披萨
type PepperoniSimplePizza struct {
	SimplePizza
}
