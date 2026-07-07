package network

import (
	"fmt"
	"time"
)

func StartMonitor() {
	go func() {
		for {
			time.Sleep(5 * time.Second)
			Mu.Lock()
			last := LastHeartbeat
			Mu.Unlock()
			if last.IsZero() {
				continue
			}
			if time.Since(last) > 15*time.Second {
				fmt.Println("Heartbeat Timeout")
				return
			}
		}

	}()
}
