package decorator_pattern

//定义火锅基本接口
type HotPotIn interface {
	GetDescription() string //获取说明
	Prices() float64        //价格
}

//定义汤底超类
type SoupBottom struct {
	Description string  //说明
	Price       float64 //价格
}

//麻辣锅底继承超类
type SpicyHotSoupBottom struct {
	SoupBottom
}

//排骨锅底继承超类
type SpareribsSoupBottom struct {
	SoupBottom
}

//new一个麻辣锅底
func NewSpicyHotSoupBottom(description string, price float64) HotPotIn {
	return &SpicyHotSoupBottom{
		SoupBottom{
			Description: description,
			Price:       price,
		},
	}
}

//new一个排骨汤底
func NewSpareribsSoupBottom(description string, price float64) HotPotIn {
	return &SpareribsSoupBottom{
		SoupBottom{
			Description: description,
			Price:       price,
		},
	}
}

func (soup SoupBottom) GetDescription() string {
	return soup.Description
}

func (soup SoupBottom) Prices() float64 {
	return soup.Price
}
