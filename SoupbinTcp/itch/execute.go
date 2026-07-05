package itch

import (
	"fmt"
	"soupbintcp/engine"
	"soupbintcp/parser"
)

func ExecuteOrder(packet []byte, eng *engine.Engine) {
	exec, err := parser.ParseExecute(packet)
	if err != nil {
		fmt.Println("Executed Errorf")
		return
	}
	eng.HandleExec(exec)

}
