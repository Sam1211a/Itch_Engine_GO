package parser

import (
	"encoding/binary"
	"fmt"
	"soupbintcp/model"
)

func ParseTrade(packet []byte) (*model.Trade, error) {
	if len(packet) < 37 {
		return nil, fmt.Errorf("Trade packet invalid")
	}
	tr := &model.Trade{}
	tr.Timestamp = binary.BigEndian.Uint32(packet[1:5])
	tr.OrderID = binary.BigEndian.Uint64(packet[5:13])
	tr.Qty = binary.BigEndian.Uint64(packet[13:21])
	tr.BookID = binary.BigEndian.Uint32(packet[21:25])
	tr.Price = binary.BigEndian.Uint32(packet[25:29])
	tr.MatchNumber = binary.BigEndian.Uint64(packet[29:37])
	return tr, nil
}
