package network

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"soupbintcp/handler"
	// "github.com/redis/go-redis/v9/helper"
)

func Connect(conn net.Conn) net.Conn {

	login := handler.LoginBuilder(
		"ITFIC1",
		"password",
		"",
		"1",
	)
	fmt.Printf("% X\n", login)
	ln, err := conn.Write(login)
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
	_, err = io.ReadFull(conn, header)
	if err != nil {
		fmt.Println("Header Error:", err)
		return nil
	}

	length := binary.BigEndian.Uint16(header)
	fmt.Println("Lenght: ", length)

	body := make([]byte, length)
	_, err = io.ReadFull(conn, body)
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
		fmt.Printf("Session  : [%s]\n", session)
		fmt.Printf("sequence  : [%s]\n", sequence)
		fmt.Println("✅ Login Accepted")

	case 'J':
		fmt.Printf("❌ Login Rejected : %c\n", body[1])

	case 'H':
		fmt.Println("Server Heartbeat")

	case 'S':
		fmt.Println("Sequenced Data")

	default:
		fmt.Printf("Unknown Packet : %c\n", body[0])
	}
	return conn
}
