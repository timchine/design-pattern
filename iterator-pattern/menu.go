package iterator_pattern

//单个菜单菜品
type MenuItem struct {
	Name        string  //菜名
	Price       float64 //价格
	Description string  //说明简介
}

//迭代器接口
type Iterator interface {
	HasNext() bool
	Next() MenuItem
}

//生成菜单迭代器接口
type Menu interface {
	CreateIterator() Iterator
}
