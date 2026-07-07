package protocol

import (
	"fmt"
	"soupbintcp/engine"
	"soupbintcp/itch"
)

func HandlePacket(packet []byte, eng *engine.Engine) {
	// fmt.Printf("SoupBin Type=%c Length=%d\n", packet[0], len(packet))
	switch packet[0] {

	case 'H':
		// fmt.Println("Server Heartbeat")

	case 'S':
		// fmt.Println("Sequenced Data->")
		itch.Decode(packet[1:], eng)

	case 'U':
		fmt.Println("Unsequenced Data")

	case 'Z':
		fmt.Println("End Of Session")

	default:
		fmt.Printf("Unknown Packet %c\n", packet[0])
	}
}
