package main

import (
	"fmt"
	"net"
	"soupbintcp/network"
	"soupbintcp/protocol"
)

func main() {

	conn, err := net.Dial("tcp", "10.61.90.20:54001")
	if err != nil {
		fmt.Println("Failed tcp connection")
		return
	}
	defer conn.Close()
	fmt.Println("tcp Connected")
	network.Connect(conn)

	for {
		packet, err := protocol.ReadPacket(conn)
		if err != nil {
			fmt.Printf("Packet read error: %T : %v\n", err, err)
			return
		}
		// fmt.Printf("RX: % X\n", packet)
		protocol.HandlePacket(packet)
	}
}
