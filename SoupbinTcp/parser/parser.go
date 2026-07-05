package parser

import (
	"encoding/binary"
	"fmt"
	"soupbintcp/model"
)

func Parse(packet []byte) (*model.Order, error) {
	if len(packet) < 30 {
		fmt.Println("invalid packet")
		return nil, fmt.Errorf("invalid packet")
	}
	order := &model.Order{}
	order.OrderId = binary.BigEndian.Uint64(packet[5:13])
	order.Price = binary.BigEndian.Uint32(packet[13:17])
	order.Qty = binary.BigEndian.Uint64(packet[17:21])
	order.Side = packet[21]
	// order.Symbol = string(packet[22:30])
	return order, nil

	// fmt.Println("========== ADD ORDER ==========")
	// fmt.Println("OrderID :", order.OrderId)
	// fmt.Println("Price   :", order.Price)
	// fmt.Println("Qty     :", order.Qty)
	// fmt.Printf("Side    : %c\n", order.Side)
	// fmt.Println("Symbol  :", order.Symbol)
}
