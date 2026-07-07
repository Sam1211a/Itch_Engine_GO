package parser

import (
	"encoding/binary"
	"fmt"
	"soupbintcp/model"
)

func ParseExecute(packet []byte) (*model.ExecuteOrder, error) {

	if len(packet) < 29 {
		return nil, fmt.Errorf("invalid execute packet")
	}

	exe := &model.ExecuteOrder{}

	exe.Timestamp = binary.BigEndian.Uint32(packet[1:5])

	exe.OrderId = binary.BigEndian.Uint64(packet[5:13])

	exe.Qty = binary.BigEndian.Uint64(packet[13:21])

	exe.MatchNumber = binary.BigEndian.Uint64(packet[21:29])

	return exe, nil
}
