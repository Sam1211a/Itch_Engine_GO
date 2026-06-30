package parser

import (
	"encoding/binary"
	"fmt"
	"soupbintcp/model"
)

func ParseReplace(packet []byte) (*model.ReplaceOrder, error) {
	if len(packet) < 21 {
		return nil, fmt.Errorf("invalid replace packet")
	}
	replace := &model.ReplaceOrder{}
	replace.OrderId = binary.BigEndian.Uint64(packet[5:13])
	replace.Price = binary.BigEndian.Uint32(packet[13:17])
	replace.Qty = binary.BigEndian.Uint32(packet[17:21])
	return replace, nil
}
