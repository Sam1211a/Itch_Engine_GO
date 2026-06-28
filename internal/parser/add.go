package parser

import (
	"encoding/binary"
	"fmt"
	"go-itch/internal/model"
)

func ParseAdd(packet []byte) (*model.Order, error) {
	if len(packet) < 30 {
		// fmt.Println("invalid packet")
		return nil, fmt.Errorf("invalid packet")
	}
	order := &model.Order{}
	order.OrderId = binary.BigEndian.Uint64(packet[5:13])
	order.Price = binary.BigEndian.Uint32(packet[13:17])
	order.Qty = binary.BigEndian.Uint32(packet[17:21])
	order.Side = packet[21]
	order.Symbol = string(packet[22:30])
	return order, nil
}
