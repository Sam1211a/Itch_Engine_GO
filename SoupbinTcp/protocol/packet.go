package protocol

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"soupbintcp/itch"
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
func HandlePacket(packet []byte) {
	switch packet[0] {

	case 'H':
		fmt.Println("Server Heartbeat")

	case 'S':
		fmt.Println("Sequenced Data->")
		itch.Decode(packet[1:])
	case 'R':
		length := binary.BigEndian.Uint16(packet)
		fmt.Println("itch len :", string(packet), length)
		itch.Decode(packet)
	case 'A':
		fmt.Printf("ADD ORDER ::::::\n %X\n", packet[1:])

	case 'U':
		fmt.Println("Unsequenced Data")

	case 'Z':
		fmt.Println("End Of Session")

	default:
		fmt.Printf("Unknown Packet %c\n", packet[0])
	}
}
