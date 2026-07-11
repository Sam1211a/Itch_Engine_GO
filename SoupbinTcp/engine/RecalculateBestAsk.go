package engine

import (
	"math"
	"soupbintcp/model"
)

func (e *Engine) RecalculateBestAsk(
	market *model.Market,
) {

	bestPrice := uint32(math.MaxUint32)
	var bestQty uint64

	for price, qty := range market.AskLevels {

		if price < bestPrice {

			bestPrice = price
			bestQty = qty
		}
	}

	if len(market.AskLevels) == 0 {

		market.BestAskPrice = 0
		market.BestAskQty = 0
		return
	}

	market.BestAskPrice = bestPrice
	market.BestAskQty = bestQty
}
