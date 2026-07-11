package engine

import (
	"fmt"
	"soupbintcp/model"
)

func (e *Engine) HandleAdd(order *model.Order) {
	if order.OrderId == 0 {
		fmt.Printf(
			"Reference Price Update Book=%d Price=%d\n",
			order.BookID,
			order.Price,
		)
		return
	}

	sec, ok := e.Securities[order.BookID]
	if ok {

		order.Symbol = sec.SecurityCode
	}
	if order.Side == 'B' {
		e.UpdateBid(order.BookID, order.Price, order.Qty)
	} else {
		e.UpdateAsk(order.BookID, order.Price, order.Qty)
	}
	e.Book.AddOrder(order)
	e.PrintMarket(order.BookID)
	// fmt.Println("Engine -> AddOrder")
}
