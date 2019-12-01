package decorator_pattern

//定义配菜超类
type Garnish struct {
	soupBottom  HotPotIn
	Description string  //说明
	Price       float64 //价格
}

//芹菜
type CeleryGarnish struct {
	Garnish
}

//白菜
type CabbageGarnish struct {
	Garnish
}

//牛肉
type BeefGarnish struct {
	Garnish
}

//青菜
type GreensGarnish struct {
	Garnish
}

func (garnish Garnish) GetDescription() string {
	return garnish.soupBottom.GetDescription() + "加" + garnish.Description
}

func (garnish Garnish) Prices() float64 {
	return garnish.soupBottom.Prices() + garnish.Price
}

func (garnish Garnish) AddGarnishTo(soupBottom HotPotIn) HotPotIn {
	garnish.soupBottom = soupBottom
	return garnish
}
