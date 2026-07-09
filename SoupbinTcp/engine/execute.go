package engine

import (
	"fmt"
	"soupbintcp/model"
)

func (e *Engine) HandleExec(exec *model.ExecuteOrder) {
	order, ok := e.Book.Get(exec.OrderId)
	if !ok {
		fmt.Println("Order not found:", exec.OrderId)
		return
	}
	// fmt.Printf(
	// 	"Execute Update -> OrderID=%d Qty=%d\n",
	// 	order.OrderId,
	// 	order.Qty,
	// )
	if order.Qty <= exec.Qty {
		e.Book.CancelOrder(exec.OrderId)
		fmt.Println("Order Fully Executed")
		return
	}
	order.Qty -= exec.Qty
	e.Book.Orders[order.OrderId] = order
	// fmt.Println("Order Partialy Executed")
}
