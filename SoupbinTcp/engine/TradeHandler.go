package engine

import (
	"fmt"
	"soupbintcp/model"
)

func (e *Engine) HandleTrade(trade *model.Trade) {

	fmt.Printf(
		"Trade -> OrderID=%d BookID=%d MatchNumber=%d Qty=%d Price=%d\n",
		trade.OrderID,
		trade.BookID,
		trade.MatchNumber,
		trade.Qty,
		trade.Price,
	)
}
