package model

type Order struct {
	OrderId uint64
	Price   uint32
	Qty     uint32
	Side    byte
	Symbol  string
}
