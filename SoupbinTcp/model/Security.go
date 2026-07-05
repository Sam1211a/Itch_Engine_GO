package model

type SecurityDirectory struct {
	MessageType byte
	Timestamp   uint32
	Orderbook   uint32
	PriceType   byte

	ISIN         string
	SecurityCode string
	Currency     string
	Group        string

	MinimumQuantity uint64
	QuantityTableID uint32
	PriceTableID    uint32
	PriceDecimals   uint32

	DelistingDate uint32
	DelistingTime uint32

	MarketType  byte
	CompanyID   uint32
	ListingType byte

	Sector       string
	Instrument   string
	SecurityName string

	MaturityDate uint32
}
