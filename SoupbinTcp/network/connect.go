package network

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"soupbintcp/handler"
	"soupbintcp/model"
	"strconv"
	"strings"

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
		c.Session,
		strconv.FormatUint(c.Sequence, 10),
	)
	fmt.Println("Session:- ", c.Session)
	fmt.Println("Sequence:- ", strconv.FormatUint(c.Sequence, 10))
	// fmt.Printf("% X\n", login)
	ln, err := c.Conn.Write(login)
	if err != nil {
		fmt.Println("login Failed")
		c.Close()
		return err
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
		return fmt.Errorf("Header Error:")
	}

	length := binary.BigEndian.Uint16(header)
	fmt.Println("Lenght: ", length)

	body := make([]byte, length)
	_, err = io.ReadFull(c.Conn, body)
	if err != nil {
		return fmt.Errorf("Body Error:")

	}

	fmt.Printf("Body = % X\n", body)
	fmt.Printf("Packet Type = %c\n", body[0])
	switch body[0] {

	case 'A':
		session := strings.TrimSpace(string(body[1:11]))
		sequence := strings.TrimSpace(string(body[11:31]))
		c.Session = session
		seq, err := strconv.ParseUint(sequence, 10, 64)
		if err == nil {
			c.Sequence = seq
		}
		model.Mu.Lock()
		model.LastHeartbeat = time.Now()
		model.Mu.Unlock()
		fmt.Printf("Session  : [%s]\n", c.Session)
		fmt.Printf("sequence  : [%d]\n", c.Sequence)
		fmt.Println("✅ Login Accepted")

	case 'J':
		fmt.Printf("❌ Login Rejected : %c\n", body[1])
		return fmt.Errorf("Login Rejected")

	case 'H':
		// fmt.Println("Server Heartbeat")
		model.Mu.Lock()
		model.LastHeartbeat = time.Now()
		model.Mu.Unlock()

	case 'S':
		fmt.Println("Sequenced Data")

	default:
		return fmt.Errorf("Unknown Packet : %c\n", body[0])
	}
	return nil
}
