package itch

import (
	"encoding/binary"
	"fmt"
)

type TimeMessege struct {
	MessegeType byte
	second      uint32
}

func DecodeTime(msg []byte) {

	t := TimeMessege{}
	t.MessegeType = msg[0]
	t.second = binary.BigEndian.Uint32(msg[1:5])
	fmt.Printf("MessageTime : %c", t.MessegeType)
	fmt.Printf("Time : %d\n", t.second)
}
