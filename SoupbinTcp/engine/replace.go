package engine

import (
	"fmt"
	"soupbintcp/model"
)

func (e *Engine) HandleReplace(rep *model.ReplaceOrder) {
	oldOrder, ok := e.Book.Get(rep.OriginalOrder)
	if !ok {
		fmt.Println("Older Order not found !")
		return
	}
	e.Book.CancelOrder(rep.OriginalOrder)
	oldOrder.OrderId = rep.NewOrder
	oldOrder.Qty = rep.Qty
	oldOrder.Price = rep.Price
	e.Book.AddOrder(oldOrder)
	fmt.Println("Engine-> ReplaceOrder")

}
