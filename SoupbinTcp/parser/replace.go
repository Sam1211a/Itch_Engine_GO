package parser

import (
	"encoding/binary"
	"fmt"
	"soupbintcp/model"
)

func ParseReplace(packet []byte) (*model.ReplaceOrder, error) {
	if len(packet) < 37 {
		return nil, fmt.Errorf("invalid replace packet")
	}
	replace := &model.ReplaceOrder{}
	replace.Timestamp = binary.BigEndian.Uint32(packet[1:5])
	replace.OriginalOrder = binary.BigEndian.Uint64(packet[5:13])
	replace.NewOrder = binary.BigEndian.Uint64(packet[13:21])
	replace.Qty = binary.BigEndian.Uint64(packet[21:29])
	replace.Price = binary.BigEndian.Uint32(packet[29:33])
	replace.Yield = binary.BigEndian.Uint32(packet[33:37])
	// fmt.Printf(
	// 	"Replace Parsed -> Old=%d New=%d Qty=%d Price=%d \n",
	// 	replace.OriginalOrder,
	// 	replace.NewOrder,
	// 	replace.Qty,
	// 	replace.Price,
	// )
	return replace, nil
}
