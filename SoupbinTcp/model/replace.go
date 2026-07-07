package model

type ReplaceOrder struct {
	Timestamp     uint32
	OriginalOrder uint64
	NewOrder      uint64
	Qty           uint64
	Price         uint32
	Yield         uint32
}
