package iterator_pattern

//定义煎饼店的菜单
type PancakeHouseMenu struct {
	menuItems []MenuItem
}

func NewPancakeHouseMenu() PancakeHouseMenu {
	var (
		p PancakeHouseMenu
	)
	p.Add("杂粮煎饼", 12, "健康美味有营养")
	p.Add("鸡蛋灌饼", 14, "没有鸡只有蛋")
	return p
}

func (p PancakeHouseMenu) CreateIterator() Iterator {
	p = NewPancakeHouseMenu()
	return &PancakeHouseMenuIterator{menuItems: p.GetMenuItems()}
}

func (p *PancakeHouseMenu) Add(name string, price float64, description string) {
	p.menuItems = append(p.menuItems, MenuItem{Name:name, Price: price, Description: description})
}

func (p *PancakeHouseMenu) GetMenuItems() []MenuItem {
	return p.menuItems
}

//定义煎饼店的菜单迭代器
type PancakeHouseMenuIterator struct {
	menuItems []MenuItem
	position int
}

func (p *PancakeHouseMenuIterator) HasNext() bool {
	if p.position >= len(p.menuItems) || p.menuItems == nil {
		return false
	} else {
		return true
	}
}

func (p *PancakeHouseMenuIterator) Next() MenuItem {
	menuItem := p.menuItems[p.position]
	p.position += 1
	return menuItem
}



