package network

import (
	"sync"
	"time"
)

var (
	LastHeartbeat time.Time
	Mu            sync.Mutex
)
