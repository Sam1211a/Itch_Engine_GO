package parser

import (
	"encoding/binary"
	"fmt"
	"soupbintcp/model"
)

func ParseDelete(packet []byte) (*model.DeleteOrder, error) {
	if len(packet) < 13 {
		return nil, fmt.Errorf("DeleteOrder error")
	}
	del := &model.DeleteOrder{}
	del.Timestamp = binary.BigEndian.Uint32(packet[1:5])
	del.OrderId = binary.BigEndian.Uint64(packet[5:13])
	return del, nil
}
