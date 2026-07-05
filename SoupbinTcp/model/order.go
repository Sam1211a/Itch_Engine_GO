package model

type Order struct {
	OrderId   uint64
	BookID    uint32
	Side      byte
	Price     uint32
	Qty       uint64
	Timestamp uint32
	Symbol    string
}
type ExecuteOrder struct {
	OrderId uint64
	Qty     uint64
}
type ReplaceOrder struct {
	OrderId uint64
	Price   uint32
	Qty     uint32
}
