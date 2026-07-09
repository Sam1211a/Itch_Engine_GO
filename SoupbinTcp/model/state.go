package model

import (
	"sync"
	"time"
)

var (
	LastHeartbeat time.Time
	Mu            sync.Mutex
)
var CurrentTime uint32
var TradingDate time.Time

func InitTradingDate() {
	today := time.Now()

	TradingDate = time.Date(
		today.Year(),
		today.Month(),
		today.Day(),
		0,
		0,
		0,
		0,
		time.Local,
	)
}
