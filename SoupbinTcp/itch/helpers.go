package itch

import (
	"encoding/binary"
	"soupbintcp/model"
	"strings"
)

func ReadString(b []byte) string {
	return strings.TrimRight(string(b), " ")
}

func ReadUint32(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}

func ReadUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}
func DecodeTime(msg []byte) uint32 {
	// fmt.Println("Current Time:%V", msg)
	model.CurrentTime = binary.BigEndian.Uint32(msg)
	return binary.BigEndian.Uint32(msg)
}
