package component_pattern

import "fmt"

type MenuItem struct {
	MenuComponent
	name        string
	price       float64
	description string
}

func (m *MenuItem) GetName() string {
	return m.name
}

func (m *MenuItem) GetPrice() float64 {
	return m.price
}

func (m *MenuItem) GetDescription() string {
	return m.description
}

func (m *MenuItem) String() string {
	return fmt.Sprintf("name:%s, price:%f, description:%s\n", m.name, m.price, m.description)
}

func NewMenuItem(name, description string, price float64) MenuComponenter {
	return &MenuItem{name: name, description: description, price: price}
}
