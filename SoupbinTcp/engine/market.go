package engine

import "soupbintcp/model"

func (e *Engine) GetMarket(bookID uint32) (*model.Market, bool) {

	m, ok := e.Markets[bookID]

	return m, ok

}

func (e *Engine) UpdateBid(bookID uint32, price uint32, qty uint64) {
	market, ok := e.Markets[bookID]
	if !ok {
		return
	}
	market.BidLevels[price] += qty
	if price > market.BestBidPrice {
		market.BestBidPrice = price
	}
	market.BestBidQty = market.BidLevels[market.BestBidPrice]
}

func (e *Engine) UpdateAsk(bookID uint32, price uint32, qty uint64) {
	market, ok := e.Markets[bookID]
	if !ok {
		return
	}
	market.AskLevels[price] += qty
	if market.BestAskPrice == 0 || price < market.BestAskPrice {
		market.BestAskPrice = price
	}
	market.BestAskQty = market.AskLevels[market.BestAskPrice]
}
func (e *Engine) UpdateTrade(bookID uint32, price uint32, qty uint64) {
	market, ok := e.Markets[bookID]
	if !ok {
		return
	}
	market.LastPrice = price
	market.LastQty = qty
	market.Volume += qty
	market.Trades++
}

func (e *Engine) RemoveBid(bookID uint32, price uint32, qty uint64) {
	market, ok := e.Markets[bookID]
	if !ok {
		return
	}
	levelQty, ok := market.BidLevels[price]
	if !ok {
		return
	}
	if levelQty <= qty {
		delete(market.BidLevels, price)
	} else {
		market.BidLevels[price] = levelQty - qty
	}
	e.RecalculateBestBid(market)
}

func (e *Engine) RemoveAsk(bookID uint32, price uint32, qty uint64) {
	market, ok := e.Markets[bookID]
	if !ok {
		return
	}
	levelQty, ok := market.AskLevels[price]
	if !ok {
		return
	}
	if levelQty <= qty {
		delete(market.AskLevels, price)
	} else {
		market.AskLevels[price] = levelQty - qty
	}
	e.RecalculateBestAsk(market)
}
