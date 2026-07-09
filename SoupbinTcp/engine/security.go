package engine

import (
	"soupbintcp/model"
)

func (e *Engine) HandleSecurity(sec *model.SecurityDirectory) {
	e.Securities[sec.Orderbook] = *sec
	// fmt.Printf(
	// 	"Security Added : %d -> %s\n",
	// 	sec.Orderbook,
	// 	sec.SecurityCode,
	// )
}
