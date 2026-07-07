package engine

import (
	"fmt"
	"soupbintcp/model"
)

func (e *Engine) HandleTrade(trade *model.Trade) {
	fmt.Printf(
		"Trade -> OrderID=%d BookID=%d Side=%c Qty=%d Price=%d\n",
		// trade.OrderId,
		trade.BookID,
		// trade.Side,
		trade.Qty,
		trade.Price,
	)
}
