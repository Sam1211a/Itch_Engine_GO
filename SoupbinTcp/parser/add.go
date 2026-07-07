package parser

import (
	"encoding/binary"
	"fmt"
	"soupbintcp/model"
)

func ParseAdd(packet []byte) (*model.Order, error) {
	// fmt.Printf("% X\n", packet)

	if len(packet) < 34 {
		return nil, fmt.Errorf("invalid add packet")
	}

	order := &model.Order{}

	order.Timestamp = binary.BigEndian.Uint32(packet[1:5])

	order.OrderId = binary.BigEndian.Uint64(packet[5:13])

	order.Side = packet[13]

	order.Qty = binary.BigEndian.Uint64(packet[14:22])

	order.BookID = binary.BigEndian.Uint32(packet[22:26])

	order.Price = binary.BigEndian.Uint32(packet[26:30])

	return order, nil
}
