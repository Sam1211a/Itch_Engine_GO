package network

import (
	"context"
	"fmt"
	"soupbintcp/engine"
	"soupbintcp/protocol"
	"time"
)

func NewClient(eng *engine.Engine) *Client {
	return &Client{
		Eng:      eng,
		Sequence: 1,
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
	if c.Cancel != nil {
		c.Cancel()
		c.Cancel = nil
	}
	c.Ctx = nil
	if c.Conn != nil {
		fmt.Println("Closing Connection...")
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
		c.Ctx, c.Cancel = context.WithCancel(context.Background())
		go c.SendHeartbeatLoop()

		// StartMonitor()

		err = c.ReadLoop()

		fmt.Println("ReadLoop Exit :", err)

		c.Close()
		fmt.Println("Reconnecting in 5 seconds...")
		time.Sleep(5 * time.Second)

	}
}
