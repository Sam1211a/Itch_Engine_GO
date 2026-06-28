package main

import (
	"encoding/binary"
	"fmt"
	"go-itch/internal/orderbook"
	"go-itch/internal/parser"

	"net"
)

func main() {
	conn, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Println("Server running on : 9000")

	listener, err := conn.Accept()
	if err != nil {
		panic(err)
	}

	fmt.Println("Client connect: ", listener.RemoteAddr())

	book := orderbook.New()

	stream := make([]byte, 0)
	buffer := make([]byte, 4096)
	for {

		n, err := listener.Read(buffer)
		if err != nil {
			fmt.Println("read error")
			return
		}
		stream = append(stream, buffer[:n]...)
		for len(stream) >= 4 {

			length := binary.BigEndian.Uint32(stream[0:4])
			fmt.Println("length: ", length)
			if len(stream) < int(length) {
				break
			}
			packet := stream[:length]
			stream = stream[length:]
			// order, err := parser.Parse(packet)

			msgType := packet[4]
			switch msgType {
			case 'A', 'C', 'B':
				order, err := parser.ParseAdd(packet)
				if err != nil {
					fmt.Println(err)
					break
				}
				book.AddOrder(*order)
			case 'X':
				orderId, err := parser.ParseCancel(packet)
				if err != nil {
					fmt.Println(err)
					break
				}
				book.CancelOrder(orderId)
			default:
				fmt.Println("Unknown Message")
			}
			// parser.Parse(packet)
			// book.AddOrder(*order)
			book.Print()
		}
	}
	// defer listener.Close()
}
