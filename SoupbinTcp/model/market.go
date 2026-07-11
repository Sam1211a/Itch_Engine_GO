package model

type Market struct {
	BookID uint32
	Symbol string
	Name   string
	Group  string

	BidLevels map[uint32]uint64
	AskLevels map[uint32]uint64

	BestBidPrice uint32
	BestAskPrice uint32

	BestBidQty uint64
	BestAskQty uint64

	LastPrice uint32
	LastQty   uint64

	Volume uint64
	Trades uint64
}
