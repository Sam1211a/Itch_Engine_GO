package network

import (
	"fmt"
	"soupbintcp/model"
	"time"
)

func StartMonitor() {
	go func() {
		for {
			time.Sleep(5 * time.Second)
			model.Mu.Lock()
			last := model.LastHeartbeat
			model.Mu.Unlock()
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
