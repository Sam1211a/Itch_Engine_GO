package network

import (
	"encoding/binary"
)

func (c *Client) SendHeartbeat() error {
	packet := make([]byte, 3)
	binary.BigEndian.PutUint16(packet[0:2], 1)
	packet[2] = 'R'
	_, err := c.Conn.Write(packet)
	return err
}
