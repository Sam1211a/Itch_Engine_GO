package network

import (
	"fmt"
	"net"
	"soupbintcp/engine"
	"soupbintcp/protocol"
	"time"
)

type Client struct {
	Conn net.Conn
	Eng  *engine.Engine
}

func NewClient(eng *engine.Engine) *Client {
	return &Client{
		Eng: eng,
	}
}
func (c *Client) ReadLoop() error {
	for {
		packet, err := protocol.ReadPacket(c.Conn)
		if err != nil {
			return err
		}
		protocol.HandlePacket(packet, c.Eng)
	}
}

func (c *Client) Close() {
	if c.Conn != nil {
		c.Conn.Close()
		c.Conn = nil
	}
}
func (c *Client) Run() {
	for {
		err := c.Connect()
		if err != nil {
			fmt.Println("Connect Failed")

			time.Sleep(5 * time.Second)

			continue
		}
		err = c.Login()

		if err != nil {

			c.Close()

			time.Sleep(5 * time.Second)

			continue
		}
		go c.SendHeartbeat()

		StartMonitor()

		err = c.ReadLoop()

		fmt.Println("ReadLoop Exit :", err)

		c.Close()

		time.Sleep(5 * time.Second)

	}
}
