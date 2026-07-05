package model

type AddOrder struct {
	MessageType byte

	Timestamp uint32

	OrderID uint64

	Side byte

	Quantity uint64

	OrderbookID uint32

	Price uint32

	Participant uint32
}
