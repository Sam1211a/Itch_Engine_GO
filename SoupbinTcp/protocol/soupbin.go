package protocol

import (
	"encoding/binary"
	"io"
	"net"
)

func ReadPacket(conn net.Conn) ([]byte, error) {

	header := make([]byte, 2)

	_, err := io.ReadFull(conn, header)
	if err != nil {
		return nil, err
	}

	length := binary.BigEndian.Uint16(header)

	body := make([]byte, length)

	_, err = io.ReadFull(conn, body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
