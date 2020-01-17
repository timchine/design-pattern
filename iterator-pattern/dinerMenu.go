package iterator_pattern

import (
	"container/list"
)

//小餐馆菜单
type DinerMenu struct {
	menuItems list.List
}


func NewDinerMenu() DinerMenu {
	var (
		d DinerMenu
	)
	d.Add("肉包子", 1, "满满的都是肉")
	d.Add("青菜香菇包", 1, "菜才最美味")
	return d
}

func (d *DinerMenu) Add(name string, price float64, description string) {
	d.menuItems.PushBack(MenuItem{Name:name, Price:price, Description:description})
}

func (d DinerMenu) CreateIterator() Iterator {
	d = NewDinerMenu()
	return &DinerMenuIterator{d.menuItems.Front(), true}
}

//定义小餐馆的迭代器
type DinerMenuIterator struct {
	cur *list.Element
	end bool
}

func (d *DinerMenuIterator) HasNext() bool {
	return d.end
}

func (d *DinerMenuIterator) Next() MenuItem {
	menuItem := d.cur.Value.(MenuItem)
	d.cur = d.cur.Next()
	if d.cur == nil {
		d.end = false
	}
	return menuItem
}