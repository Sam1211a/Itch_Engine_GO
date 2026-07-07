package itch

import (
	"soupbintcp/engine"
	"soupbintcp/parser"
)

func DeleteOrder(msg []byte, eng *engine.Engine) {
	del, err := parser.ParseDelete(msg)
	if err != nil {
		return
	}
	eng.HandleDelete(del)
}
