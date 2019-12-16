package facade_pattern

import "fmt"

type GoodsService struct {
	goods Goods
}

type Goods struct {
	Id int
	Name string
	Price int
}

func (goodsService *GoodsService) GetGoodsById() {
	fmt.Println("通过商品id查询，库存是否大于0")
}

func (goodsService GoodsService) SumGoodsPrice() {
	fmt.Println("计算商品的价格")
}