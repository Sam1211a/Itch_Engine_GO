package parser

import (
	"encoding/binary"
	"fmt"
)

func ParseCancel(packet []byte) (uint64, error) {
	if len(packet) < 13 {
		fmt.Println("invalid packet")
		return 0, fmt.Errorf("invalid packet")
	}
	OrderID := binary.BigEndian.Uint64(packet[5:13])
	return OrderID, nil
}
