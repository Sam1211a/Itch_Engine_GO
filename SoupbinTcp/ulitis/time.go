package ulitis

import (
	"soupbintcp/model"
	"time"
)

func ITCHTime(ns uint32) time.Time {
	// fmt.Println("CurrentSecond:", model.CurrentTime)
	// fmt.Println("Nano:", ns)
	t := model.TradingDate.
		Add(time.Duration(model.CurrentTime) * time.Second).
		Add(time.Duration(ns) * time.Nanosecond)
	return t
}
func ITCHTimeString(ns uint32) string {

	return ITCHTime(ns).Format(
		"2006-01-02 03:04:05.000000000",
	)
}
