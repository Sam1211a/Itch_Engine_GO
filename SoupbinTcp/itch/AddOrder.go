package itch

import (
	"soupbintcp/engine"
	"soupbintcp/parser"
)

func AddOrder(msg []byte, eng *engine.Engine) {
	// book := orderbook.New()
	order, err := parser.ParseAdd(msg)
	if err != nil {
		return
	}
	eng.HandleAdd(order)
	// book.Print()

	// fmt.Println("============== ADD ORDER ==============")
	// fmt.Println("Timestamp :", a.Timestamp)
	// fmt.Println("OrderID   :", a.OrderID)
	// fmt.Printf("Side      : %c\n", a.Side)
	// fmt.Println("Qty       :", a.Quantity)
	// fmt.Println("BookID    :", a.OrderbookID)
	// fmt.Println("Price     :", a.Price)
	// fmt.Println("Participant:", a.Participant)
}
