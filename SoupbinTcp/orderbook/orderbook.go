package orderbook

import (
	"fmt"
	"soupbintcp/model"
)

type Orderbook struct {
	Orders map[uint64]model.Order
}

func New() *Orderbook {
	return &Orderbook{
		Orders: make(map[uint64]model.Order),
	}
}

func (ob *Orderbook) AddOrder(order model.Order) {
	ob.Orders[order.OrderId] = order
	fmt.Println("Order added :", order.OrderId)
}

func (ob *Orderbook) Get(OrderID uint64) (model.Order, bool) {
	order, ok := ob.Orders[OrderID]
	return order, ok
}
func (ob *Orderbook) CancelOrder(OrderID uint64) {
	fmt.Println("Order Cancel: ", ob.Orders[OrderID].OrderId)
	delete(ob.Orders, OrderID)
}
func (ob *Orderbook) Print() {
	fmt.Println("----------- ORDER BOOK -----------")

	for id, order := range ob.Orders {
		fmt.Printf(
			"ID=%d Symbol=%s Price=%d Qty=%d Side=%c\n",
			id,
			order.Price,
			order.Qty,
			order.Side,
		)
	}

	fmt.Println("----------------------------------")
}
