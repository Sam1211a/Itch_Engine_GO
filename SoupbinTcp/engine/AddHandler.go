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
	sec := e.Securities[order.BookID]
	order.Symbol = sec.SecurityCode
	e.Book.AddOrder(*order)
	// fmt.Println("Engine -> AddOrder")
}
