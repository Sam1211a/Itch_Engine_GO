package engine

import (
	"fmt"
	"soupbintcp/model"
)

func (e *Engine) HandleAdd(order *model.Order) {
	sec := e.Securities[order.BookID]
	order.Symbol = sec.SecurityCode
	e.Book.AddOrder(*order)
	fmt.Println("Engine -> AddOrder")
}
