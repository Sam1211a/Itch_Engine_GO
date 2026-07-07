package model

type Trade struct {
	Timestamp uint32

	OrderID uint64

	Qty uint64

	BookID uint32

	Price uint32

	MatchNumber uint64
}
