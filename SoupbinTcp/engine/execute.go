package engine

import (
	"fmt"

	"soupbintcp/model"
	"soupbintcp/ulitis"
)

func (e *Engine) HandleExec(exec *model.ExecuteOrder) {
	order, ok := e.Book.Get(exec.OrderId)
	if !ok {
		fmt.Println("Order not found:", exec.OrderId)
		return
	}
	market, ok := e.GetMarket(order.BookID)
	if !ok {
		return
	}
	fmt.Printf(
		"Execute Update -> OrderID = %d Qty = %d Timestamp = %s\n",
		order.OrderId,
		order.Qty,
		ulitis.ITCHTimeString(exec.Timestamp),
	)
	market.LastPrice = order.Price
	market.LastQty = exec.Qty
	market.Volume += exec.Qty
	market.Trades++
	if order.Side == 'B' {
		e.RemoveBid(order.BookID, order.Price, exec.Qty)
	} else {
		e.RemoveAsk(order.BookID, order.Price, exec.Qty)
	}
	if order.Qty <= exec.Qty {
		e.Book.CancelOrder(exec.OrderId)
		fmt.Println("Order Fully Executed")
		return
	}
	order.Qty -= exec.Qty
	e.Book.Orders[order.OrderId] = order
	// fmt.Println("Order Partialy Executed")
}
