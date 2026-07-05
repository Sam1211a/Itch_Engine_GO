package parser

import (
	"encoding/binary"
	"fmt"
	"soupbintcp/model"
)

func ParseExecute(packet []byte) (*model.ExecuteOrder, error) {
	if len(packet) < 17 {
		return nil, fmt.Errorf("invalid Executed packet")
	}
	order := &model.ExecuteOrder{}
	order.OrderId = binary.BigEndian.Uint64(packet[5:13])
	order.Qty = binary.BigEndian.Uint64(packet[13:17])
	return order, nil
}
