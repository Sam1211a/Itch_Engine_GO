package engine

import "soupbintcp/model"

func (e *Engine) RecalculateBestBid(market *model.Market) {
	var bestprice uint32
	var bestqty uint64
	for price, qty := range market.BidLevels {
		if qty > bestqty {
			bestprice = price
			bestqty = qty
		}
	}
	market.BestBidPrice = bestprice
	market.BestBidQty = bestqty
}
