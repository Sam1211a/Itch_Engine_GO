package engine

import (
	"soupbintcp/model"
)

func (e *Engine) HandleSecurity(sec *model.SecurityDirectory) {
	e.Securities[sec.Orderbook] = sec
	e.Markets[sec.Orderbook] = &model.Market{
		BookID:    sec.Orderbook,
		Symbol:    sec.SecurityCode,
		Name:      sec.SecurityName,
		Group:     sec.Group,
		BidLevels: make(map[uint32]uint64),
		AskLevels: make(map[uint32]uint64),
	}
	// fmt.Printf(
	// 	"Security Added : %d -> %s\n",
	// 	sec.Orderbook,
	// 	sec.SecurityCode,
	// )
}
