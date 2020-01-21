package component_pattern

import "fmt"

type MenuComponenter interface {
	Add(component MenuComponenter)
	Remove(component MenuComponenter)
	GetChild(i int) MenuComponenter
	GetName() string
	GetDescription() string
	GetPrice() float64
	fmt.Stringer
}

type MenuComponent struct {

}

func (m MenuComponent) Add(component MenuComponenter) {

}

func (m MenuComponent) Remove(component MenuComponenter) {

}

func (m MenuComponent) GetChild(i int) MenuComponenter {
	return nil
}

func (m MenuComponent) GetName() string {
	return ""
}

func (m MenuComponent) GetDescription() string {
	return ""
}

func (m MenuComponent) GetPrice() float64 {
	return 0
}

func (m MenuComponent) String() string {
	return fmt.Sprint(m)
}

