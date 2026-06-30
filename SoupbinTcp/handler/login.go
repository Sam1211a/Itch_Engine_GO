package handler

import (
	"bytes"
	"encoding/binary"
)

// fixed left-justifies s in a size-byte field, space-padded on the right.
// Used for the alphanumeric login fields.
func fixed(s string, size int) []byte {
	b := make([]byte, size)

	copy(b, []byte(s))

	for i := len(s); i < size; i++ {
		b[i] = ' '
	}

	return b
}

// fixedRight right-justifies s in a size-byte field, space-padded on the
// left. SoupBinTCP numeric fields (the requested sequence number) require
// this layout.
func fixedRight(s string, size int) []byte {
	b := make([]byte, size)

	for i := range b {
		b[i] = ' '
	}
	if len(s) > size {
		s = s[len(s)-size:]
	}
	copy(b[size-len(s):], []byte(s))

	return b
}

func LoginBuilder(
	user string,
	pass string,
	session string,
	seq string,
) []byte {

	buf := new(bytes.Buffer)

	// Reserve 2 bytes for length
	binary.Write(buf, binary.BigEndian, uint16(0))

	// Packet Type
	buf.WriteByte('L')

	// Username 6 bytes
	buf.Write(fixed(user, 6))

	// Password 10 bytes
	buf.Write(fixed(pass, 10))

	// Session 10 bytes
	buf.Write(fixed(session, 10))

	// Sequence 20 bytes (numeric, right-justified)
	buf.Write(fixedRight(seq, 20))

	packet := buf.Bytes()

	// Length = bytes after length field
	binary.BigEndian.PutUint16(
		packet[0:2],
		uint16(len(packet)-2),
	)

	return packet
}
