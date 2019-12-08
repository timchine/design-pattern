package factory_pattern

//创建对象的接口
type PizzaStoryCreate interface {
	CreatePizza(t PizzaType) IPizza
}

//店面超类
type PizzaStory struct {
	PizzaStoryCreate
}

//预定一个披萨
func (pizzaStory PizzaStory) OrderPizza(t PizzaType) IPizza {
	if pizzaStory.PizzaStoryCreate == nil {
		return nil
	}
	pizza := pizzaStory.CreatePizza(t)
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}

func NewPizzaStore(create PizzaStoryCreate) PizzaStory {
	return PizzaStory{PizzaStoryCreate: create}
}

//纽约分店
type NYPizzaStory struct {

}

func (N NYPizzaStory) CreatePizza(t PizzaType) IPizza {
	switch t {
	case CheesePizzaType:
		return NewNyCheesePizza()
	case ClamPizzaType:
		return NewNyClamPizza()
	case PepperoniPizzaType:
		return NewNyPepperoniPizza()
	}
	return nil
}

type ChinaPizzaStory struct {

}

func (c ChinaPizzaStory) CreatePizza(t PizzaType) IPizza {
	switch t {
	case CheesePizzaType:
		return NewChinaCheesePizza()
	case ClamPizzaType:
		return NewChinaClamPizza()
	case PepperoniPizzaType:
		return NewChinaPepperoniPizza()
	}
	return nil
}