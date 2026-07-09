package main

import (
	"soupbintcp/engine"
	"soupbintcp/network"
)

func main() {

	eng := engine.New()
	client := network.NewClient(eng)

	client.Run()
	// conn, err := net.Dial("tcp", network.Host)
	// if err != nil {
	// 	fmt.Println("Failed tcp connection")
	// 	return
	// }
	// defer conn.Close()
	// fmt.Println("tcp Connected")
	// network.Connect(conn)
	// network.StartMonitor()
	// go func() {
	// 	ticker := time.NewTicker(time.Second)
	// 	defer ticker.Stop()
	// 	for range ticker.C {
	// 		err := network.SendHeartbeat(conn)
	// 		if err != nil {
	// 			return
	// 		}
	// 	}

	// }()
	// eng := engine.New()
	// for {
	// 	packet, err := protocol.ReadPacket(conn)
	// 	if err != nil {
	// 		fmt.Printf("Packet read error: %T : %v\n", err, err)
	// 		break
	// 	}
	// 	// fmt.Printf("RX: % X\n", packet)
	// 	protocol.HandlePacket(packet, eng)
	// }
	// eng.PrintBook()
}
