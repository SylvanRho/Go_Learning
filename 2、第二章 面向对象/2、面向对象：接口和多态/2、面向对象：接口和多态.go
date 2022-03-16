package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println()
	iPhone := Phone{
		name:     "iPhone",
		quantity: 1,
		price:    8000,
	}
	earphones := FreeGift{
		name:     "耳机",
		quantity: 1,
		price:    200,
	}
	good := []Good{iPhone, earphones}
	allPrice := calculateAllPrice(good)
	fmt.Print("您购买的总金额为：", allPrice)
}

// 定义一个接口
type Good interface {
	//结账
	settleAccount() int
	//订单信息
	orderInfo() string
}

//定义一个手机的对象
type Phone struct {
	name            string
	quantity, price int
}

//手机实现结账接口方法
func (phone Phone) settleAccount() int {
	return phone.quantity * phone.price
}

//手机实现订单信息方法
func (phone Phone) orderInfo() string {
	return "您要购买" + strconv.Itoa(phone.quantity) + "个" +
		phone.name + "共计" + strconv.Itoa(phone.settleAccount()) + "元"
}

//定义一个礼物的对象
type FreeGift struct {
	name            string
	quantity, price int
}

//礼物实现结账接口方法
func (gift FreeGift) settleAccount() int {
	//因为免费 所以返回0
	return 0
}

//礼物实现订单信息接口方法
func (gift FreeGift) orderInfo() string {
	return "您要购买" + strconv.Itoa(gift.quantity) + "个" +
		gift.name + "共计" + strconv.Itoa(gift.settleAccount()) + "元"
}

//计算所有价格
func calculateAllPrice(good []Good) int {
	var allPrice int
	for _, good := range good {
		fmt.Println(good.orderInfo())
		allPrice += good.settleAccount()
	}
	return allPrice
}
