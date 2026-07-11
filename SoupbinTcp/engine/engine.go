package engine

import (
	"fmt"
	"soupbintcp/model"
	"soupbintcp/orderbook"
)

type Engine struct {
	Book       *orderbook.Orderbook
	Securities map[uint32]*model.SecurityDirectory
	Markets    map[uint32]*model.Market
}

func New() *Engine {
	return &Engine{
		Book:       orderbook.New(),
		Securities: make(map[uint32]*model.SecurityDirectory),
		Markets:    make(map[uint32]*model.Market),
	}
}

func (e *Engine) PrintBook() {
	e.Book.Print()
}

func (e *Engine) PrintMarket(bookID uint32) {

	market, ok := e.Markets[bookID]
	if !ok {
		fmt.Println("Market Not Found")
		return
	}

	fmt.Println("======================================")
	fmt.Printf("BookID : %d\n", market.BookID)
	fmt.Printf("Symbol : %s\n", market.Symbol)

	fmt.Println("----------- BID -----------")
	fmt.Printf("Best Bid Price : %d\n", market.BestBidPrice)
	fmt.Printf("Best Bid Qty   : %d\n", market.BestBidQty)

	fmt.Println("Bid Levels:")
	for price, qty := range market.BidLevels {
		fmt.Printf("  %d -> %d\n", price, qty)
	}

	fmt.Println("----------- ASK -----------")
	fmt.Printf("Best Ask Price : %d\n", market.BestAskPrice)
	fmt.Printf("Best Ask Qty   : %d\n", market.BestAskQty)

	fmt.Println("Ask Levels:")
	for price, qty := range market.AskLevels {
		fmt.Printf("  %d -> %d\n", price, qty)
	}

	fmt.Println("----------- TRADE -----------")
	fmt.Printf("Last Price : %d\n", market.LastPrice)
	fmt.Printf("Last Qty   : %d\n", market.LastQty)
	fmt.Printf("Volume     : %d\n", market.Volume)
	fmt.Printf("Trades     : %d\n", market.Trades)

	fmt.Println("======================================")
}
