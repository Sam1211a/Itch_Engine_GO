package itch

import (
	"encoding/binary"

	"soupbintcp/model"
)

func DecodeSecurityDirectory(msg []byte) *model.SecurityDirectory {
	s := &model.SecurityDirectory{}
	s.MessageType = msg[0]
	s.Timestamp = DecodeTime(msg[1:5])
	// DecodeTime(msg[0:5])
	s.Orderbook = binary.BigEndian.Uint32(msg[5:9])

	s.PriceType = msg[9]

	s.ISIN = ReadString(msg[10:22])
	s.SecurityCode = ReadString(msg[22:34])
	s.Currency = ReadString(msg[34:37])
	s.Group = ReadString(msg[37:45])

	s.MinimumQuantity = ReadUint64(msg[45:53])
	s.QuantityTableID = ReadUint32(msg[53:57])
	s.PriceTableID = ReadUint32(msg[57:61])
	s.PriceDecimals = ReadUint32(msg[61:65])

	s.DelistingDate = ReadUint32(msg[65:69])
	s.DelistingTime = ReadUint32(msg[69:73])

	s.MarketType = msg[73]

	s.CompanyID = ReadUint32(msg[74:78])

	s.ListingType = msg[78]

	s.Sector = ReadString(msg[79:91])
	s.Instrument = ReadString(msg[91:103])

	s.SecurityName = ReadString(msg[103:163])

	s.MaturityDate = ReadUint32(msg[163:167])

	// fmt.Printf("Timestamp: %+v\n", msg[1:5])
	// fmt.Printf("Message Type : %c\n", s.MessageType)
	// fmt.Printf("Timestamp    : %d\n", s.Timestamp)
	// fmt.Printf("Orderbook    : %d\n", s.Orderbook)
	// fmt.Printf("Price Type   : %c\n", s.PriceType)
	// fmt.Printf("ISIN         : %s\n", s.ISIN)
	// fmt.Printf("SecurityCode : %s\n", s.SecurityCode)
	// fmt.Printf("Currency :%s\n", s.Currency)
	// fmt.Printf("Group    :%s\n", s.Group)
	// fmt.Printf("MinimumQuantity: %d\n", s.MinimumQuantity)
	// fmt.Printf("QuantityTableID: %d\n", s.QuantityTableID)
	// fmt.Printf("Market Type  : %c\n", s.MarketType)
	// fmt.Printf("Company ID   : %d\n", s.CompanyID)
	// fmt.Printf("Listing Type : %c\n", s.ListingType)
	// fmt.Printf("Security Name: %s\n", s.SecurityName)

	// fmt.Println("Timestamp :", s.Timestamp)
	return s

}
