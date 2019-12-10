package factory_pattern

type PizzaIngredientFactory interface {
	CreateDough() Dough
	CreateSauce() Sauce
}

//面团
type Dough interface{}

//酱
type Sauce interface{}

//纽约的原料工厂
type NyPizzaIngredientFactory struct {
}

type NyDough struct {
	Dough
	Name string
}

type NySauce struct {
	Sauce
	Name string
}

func (n NyPizzaIngredientFactory) CreateDough() Dough {
	return &NyDough{Name: "纽约的面团"}
}

func (n NyPizzaIngredientFactory) CreateSauce() Sauce {
	return &NySauce{
		Name: "纽约的酱汁",
	}
}

//中国的原料工厂
type ChinaPizzaIngredientFactory struct {
}

type ChinaDough struct {
	Dough
	Name string
}

type ChinaSauce struct {
	Sauce
	Name string
}

func (c ChinaPizzaIngredientFactory) CreateDough() Dough {
	return &ChinaDough{Name: "中国的面团"}
}

func (c ChinaPizzaIngredientFactory) CreateSauce() Sauce {
	return &ChinaSauce{Name:"中国的酱汁"}
}
