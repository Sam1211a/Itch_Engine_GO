package itch

import (
	"encoding/binary"
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
	return binary.BigEndian.Uint32(msg)
}
