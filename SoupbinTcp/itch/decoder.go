package itch

import (
	"encoding/binary"
	"fmt"
	"soupbintcp/model"
	"strings"
)

func ReadString(b []byte) string {
	return strings.TrimRight(string(b), " ")
}
func Read_uint_32(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}

func DecodeSecurityDirectory(msg []byte) {
	s := model.SecurityDirectory{}
	s.MessageType = msg[0]
	s.Timestamp = binary.BigEndian.Uint32(msg[1:5])
	s.Orderbook = binary.BigEndian.Uint32(msg[5:9])

	s.PriceType = msg[9]

	s.ISIN = strings.TrimSpace(string(msg[10:22]))
	s.SecurityCode = strings.TrimSpace(string(msg[22:34]))

	s.Currency = strings.TrimSpace(string(msg[34:37]))
	s.Group = strings.TrimSpace(string(msg[37:45]))

	s.MinimumQuantity = binary.BigEndian.Uint64(msg[45:53])

	s.QuantityTableID = binary.BigEndian.Uint32(msg[53:57])
	s.PriceTableID = binary.BigEndian.Uint32(msg[57:61])
	s.PriceDecimals = binary.BigEndian.Uint32(msg[61:65])

	s.DelistingDate = binary.BigEndian.Uint32(msg[65:69])
	s.DelistingTime = binary.BigEndian.Uint32(msg[69:73])

	s.MarketType = msg[73]

	s.CompanyID = binary.BigEndian.Uint32(msg[74:78])

	s.ListingType = msg[78]

	s.Sector = strings.TrimSpace(string(msg[79:91]))
	s.Instrument = strings.TrimSpace(string(msg[91:103]))

	s.SecurityName = strings.TrimSpace(string(msg[103:163]))

	s.MaturityDate = binary.BigEndian.Uint32(msg[163:167])

	// fmt.Printf("Timestamp: %+v\n", msg[1:5])
	fmt.Printf("Message Type : %c\n", s.MessageType)
	fmt.Printf("Timestamp    : %d\n", s.Timestamp)
	fmt.Printf("Orderbook    : %d\n", s.Orderbook)
	fmt.Printf("Price Type   : %c\n", s.PriceType)
	fmt.Printf("ISIN         : %s\n", s.ISIN)
	fmt.Printf("SecurityCode : %s\n", s.SecurityCode)
	fmt.Printf("Currency :%c\n", s.Currency)
	fmt.Printf("Group    :%s\n", s.Group)
	fmt.Printf("MinimumQuantity: %d\n", s.MinimumQuantity)
	fmt.Printf("QuantityTableID: %d\n", s.QuantityTableID)
	fmt.Printf("Market Type  : %c\n", s.MarketType)
	fmt.Printf("Company ID   : %d\n", s.CompanyID)
	fmt.Printf("Listing Type : %c\n", s.ListingType)
	fmt.Printf("Security Name: %s\n", s.SecurityName)

	// fmt.Println("Timestamp :", s.Timestamp)

	// fmt.Printf("Raw Length : %d\n", len(msg))
}

func Decode(packet []byte) {

	if len(packet) == 0 {
		return
	}
	switch packet[0] {

	case 'R':
		DecodeSecurityDirectory(packet)

	case 'A':
		fmt.Println("Add Order")

	case 'E':
		fmt.Println("Order Executed")

	case 'P':
		fmt.Println("Trade")

	default:
		fmt.Printf("Unknown ITCH Message: %c\n", packet[0])
	}
}
