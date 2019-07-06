package alphaVantage

import (
	"time"
)

type rawError struct {
	Note string `json:"Note"`
}


type Metadata struct {
	Information string
	Symbol string
	LastRefreshed time.Time
	Interval Interval
	Size Size
}


