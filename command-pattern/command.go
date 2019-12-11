package command_pattern

import (
	"fmt"
	"time"
)

//定义队列命令接口
type Command interface {
	Execute()
}

//定义商品
var Goods = make([]int, 10)

//下单的队列
type QueueOrder struct {
	Order //订单
}

type Order struct {
	GoodsId string //商品的id
	BuyNum  int    //商品的数量
	OrderNo string //订单号
}

func (order *Order) CreateOrder() {
	order.OrderNo = time.Now().Format("20060102150405")
}

func (q QueueOrder) Execute() {
	//生成订单
	q.Order.CreateOrder()
	if q.BuyNum > 0 && len(Goods) >= q.BuyNum {
		//减少库存
		Goods = Goods[q.BuyNum:]
		//将订单入库
		fmt.Println("我已经入库")
	} else {
		fmt.Println("我已经卖完了，下次再来")
	}
}
