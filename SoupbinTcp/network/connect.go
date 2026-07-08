package network

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"soupbintcp/handler"
	"soupbintcp/model"
	"strconv"

	"time"
	// "github.com/redis/go-redis/v9/helper"
)

func (c *Client) Connect() error {
	conn, err := net.Dial("tcp", Host)
	if err != nil {
		return fmt.Errorf("Failed tcp connection")
	}
	c.Conn = conn
	fmt.Println("tcp Connected")
	return nil
}

func (c *Client) Login() error {
	// c.Sequence = 1
	login := handler.LoginBuilder(
		Username,
		Password,
		Session,
		strconv.FormatUint(c.Sequence, 10),
	)
	fmt.Printf("% X\n", login)
	ln, err := c.Conn.Write(login)
	if err != nil {
		fmt.Println("login Failed")
		return nil
	}
	fmt.Println("Bytes sent: ", ln)
	fmt.Printf("Username: [%s]\n", login[3:9])
	fmt.Printf("Password: [%s]\n", login[9:19])
	fmt.Printf("Session : [%s]\n", login[19:29])
	fmt.Printf("Sequence: [%s]\n", login[29:49])
	fmt.Println("Login request sent")
	header := make([]byte, 2)
	_, err = io.ReadFull(c.Conn, header)
	if err != nil {
		fmt.Println("Header Error:", err)
		return nil
	}

	length := binary.BigEndian.Uint16(header)
	fmt.Println("Lenght: ", length)

	body := make([]byte, length)
	_, err = io.ReadFull(c.Conn, body)
	if err != nil {
		fmt.Println("Body Error:", err)
		return nil
	}

	fmt.Printf("Body = % X\n", body)
	fmt.Printf("Packet Type = %c\n", body[0])
	switch body[0] {

	case 'A':
		session := string(body[1:11])
		sequence := string(body[11:31])
		model.Mu.Lock()
		model.LastHeartbeat = time.Now()
		model.Mu.Unlock()
		fmt.Printf("Session  : [%s]\n", session)
		fmt.Printf("sequence  : [%s]\n", sequence)
		fmt.Println("✅ Login Accepted")

	case 'J':
		fmt.Printf("❌ Login Rejected : %c\n", body[1])

	case 'H':
		fmt.Println("Server Heartbeat")
		model.Mu.Lock()
		model.LastHeartbeat = time.Now()
		model.Mu.Unlock()

	case 'S':
		fmt.Println("Sequenced Data")

	default:
		fmt.Printf("Unknown Packet : %c\n", body[0])
	}
	return nil
}
