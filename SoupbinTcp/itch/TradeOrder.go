package itch

import (
	"soupbintcp/engine"
	"soupbintcp/parser"
)

func TradeOrder(packet []byte, eng *engine.Engine) {
	trade, err := parser.ParseTrade(packet)
	if err != nil {
		return
	}
	eng.HandleTrade(trade)

}
