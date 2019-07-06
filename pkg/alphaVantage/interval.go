package alphaVantage

type Interval string
const (
	INTERVAL_1 Interval = "1min"
	INTERVAL_5 Interval = "5min"
	INTERVAL_15 Interval = "5min"
	INTERVAL_30 Interval = "30min"
	INTERVAL_60 Interval = "360min"
)

func intervalFromString(s string) Interval {
	switch s {
	case "1min": return INTERVAL_1
	case "5min": return INTERVAL_5
	case "15min": return INTERVAL_15
	case "30min": return INTERVAL_30
	case "60min": return INTERVAL_60
	default: return INTERVAL_60
	}
}

