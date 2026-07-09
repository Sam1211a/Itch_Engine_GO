package network

import (
	"encoding/binary"
	"fmt"
	"time"
)

func (c *Client) SendHeartbeat() error {
	packet := make([]byte, 3)
	binary.BigEndian.PutUint16(packet[0:2], 1)
	packet[2] = 'R'
	_, err := c.Conn.Write(packet)
	return err
}

func (c *Client) SendHeartbeatLoop() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-c.Ctx.Done():
			fmt.Println("Heartbeat Loop Exit")
			c.Close()
			return
		case <-ticker.C:
			err := c.SendHeartbeat()
			if err != nil {
				fmt.Println("SendHeartbeat Failed")
				c.Close()
				return
			}
		}
	}
}
