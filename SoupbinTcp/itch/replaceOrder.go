package itch

import (
	"soupbintcp/engine"
	"soupbintcp/parser"
)

func ReplaceOrder(packet []byte, eng *engine.Engine) {
	rep, err := parser.ParseReplace(packet)
	if err != nil {
		return
	}
	eng.HandleReplace(rep)
}
