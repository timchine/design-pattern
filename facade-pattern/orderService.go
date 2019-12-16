package facade_pattern

import "fmt"

type OrderService struct {
	GoodsService
	CarService
}

func (orderService OrderService) CreateOrder() {
	fmt.Println("生成订单")
}

func (orderService OrderService) SubOrder() {
	orderService.GetGoodsById()
	orderService.SumGoodsPrice()
	orderService.EmptyCar()
	orderService.CreateOrder()
}
