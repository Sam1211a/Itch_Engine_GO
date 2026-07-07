package itch

import (
	"fmt"
	"soupbintcp/engine"
	"soupbintcp/parser"
)

func AddOrder(msg []byte, eng *engine.Engine) {
	// fmt.Printf("ADD parser got type=%c\n", msg[0])
	// book := orderbook.New()
	order, err := parser.ParseAdd(msg)
	if err != nil {
		return
	}
	eng.HandleAdd(order)
	fmt.Printf(
		"###########################->ADD ORDER<-###################\n OrderID=%d BookID=%d Symbol=%s Side=%c Qty=%d Price=%d Timestamp=%d\n ###########################################################\n",
		order.OrderId,
		order.BookID,
		order.Symbol,
		order.Side,
		order.Qty,
		order.Price,
		order.Timestamp,
	)

}
