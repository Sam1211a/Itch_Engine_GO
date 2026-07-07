package itch

import (
	"fmt"
	"soupbintcp/engine"
)

func Decode(packet []byte, eng *engine.Engine) {
	// fmt.Printf("ITCH Type=%c Length=%d\n", packet[0], len(packet))

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
		DeleteOrder(packet, eng)
	case 'U':
		ReplaceOrder(packet, eng)

	case 'E':
		ExecuteOrder(packet, eng)

	case 'P':
		TradeOrder(packet, eng)

	default:
		fmt.Printf("Unknown ITCH Message: %c\n", packet[0])
	}
}
