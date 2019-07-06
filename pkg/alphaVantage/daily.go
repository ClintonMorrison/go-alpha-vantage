package alphaVantage

import (
"encoding/json"
	"time"
	"github.com/ClintonMorrison/goAlphaVantage/internal/parse"
)

type rawDailyMetadata struct {
	Information string `json:"1. Information"`
	Symbol string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize string `json:"4. Output Size"`
	TimeZone string `json:"5. Time Zone"`
}

func (m *rawDailyMetadata) Parse() Metadata {
	timezone, err := time.LoadLocation(m.TimeZone)
	if err != nil {
		panic(err)
	}

	lastRefreshed, err := parse.DateFromStringLocation(m.LastRefreshed, timezone)
	if err != nil {
		panic(err)
	}

	return Metadata{
		Information: m.Information,
		Symbol: m.Symbol,
		LastRefreshed: lastRefreshed,
		Interval: "DAILY",
		Size: sizeFromString(m.OutputSize),
	}
}

type rawTimeSeriesDaily struct {
	MetaData rawDailyMetadata `json:"Meta Data"`
	TimeSeriesDaily rawTimeSeries `json:"Time Series (Daily)"`
}


func (r *rawTimeSeriesDaily) Parse() *TimeSeriesDaily {
	metadata := r.MetaData.Parse()
	rawQuotes := r.TimeSeriesDaily

	return &TimeSeriesDaily{
		Metadata: metadata,
		Quotes: *rawQuotes.Parse(),
	}
}

type TimeSeriesDaily struct {
	Metadata Metadata
	Quotes TimeSeries
}

func (a *AlphaVantage) TimeSeriesDaily(symbol string, size Size) (*TimeSeriesDaily, error) {
	params := map[string]string{
		"function": "TIME_SERIES_DAILY",
		"symbol": symbol,
		"outputsize": string(size),
		"datatype": "json",
	}

	resp, err := a.request(params)
	if err != nil {
		return nil, err
	}


	raw := rawTimeSeriesDaily{}
	err = json.Unmarshal(resp.Body, &raw)
	if err != nil {
		return nil, err
	}

	return raw.Parse(), err
}
