package itch

import (
	"fmt"
	"soupbintcp/engine"
)

func Decode(packet []byte, eng *engine.Engine) {

	if len(packet) == 0 {
		return
	}
	switch packet[0] {
	case 'T':
		DecodeTime(packet)

	case 'R':
		sec := DecodeSecurityDirectory(packet)
		// if err != nil {
		// 	return
		// }
		eng.HandleSecurity(sec)
		// DecodeSecurityDirectory(packet)

	case 'A':
		AddOrder(packet, eng)
	case 'D':
		fmt.Println("DecodeDeleteorder Packet")

	case 'E':
		ExecuteOrder(packet, eng)

	case 'P':
		fmt.Println("Trade")

	default:
		fmt.Printf("Unknown ITCH Message: %c\n", packet[0])
	}
}
