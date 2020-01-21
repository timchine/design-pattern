package component_pattern

import "fmt"

type Menu struct {
	MenuComponenter
	name string
	description string
	item []MenuComponenter
}

func (m *Menu) GetName() string {
	return m.name
}

func (m *Menu) GetDescription() string {
	return m.description
}

func (m *Menu) Add(componenter MenuComponenter) {
	m.item = append(m.item, componenter)
}

func (m *Menu) Remove(componenter MenuComponenter)  {
	for k, v := range m.item {
		if componenter == v {
			m.item = append(m.item[:k], m.item[k+1:]...)
		}
	}
}

func (m *Menu) GetChild(i int) MenuComponenter{
	return m.item[i]
}

func (m *Menu) String() string {
	s := fmt.Sprintf("name:%s,description:%s-----------\n", m.name, m.description)
	for _, v := range m.item {
		s += v.String()
	}
	return s
}

func NewMenu(name, description string) MenuComponenter {
	return &Menu{name:name, description: description}
}