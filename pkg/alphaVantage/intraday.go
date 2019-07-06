package alphaVantage

import (
	"encoding/json"
	"time"
	"github.com/ClintonMorrison/goAlphaVantage/internal/parse"
)


type rawIntradayMetadata struct {
	Information string `json:"1. Information"`
	Symbol string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	Interval string `json:"4. Interval"`
	OutputSize string `json:"5. Output Size"`
	TimeZone string `json:"6. Time Zone"`
}

func (m *rawIntradayMetadata) Parse() Metadata {
	timezone, err := time.LoadLocation(m.TimeZone)
	if err != nil {
		panic(err)
	}

	lastRefreshed, err := parse.TimeFromStringLocation(m.LastRefreshed, timezone)
	if err != nil {
		panic(err)
	}

	return Metadata{
		Information: m.Information,
		Symbol: m.Symbol,
		LastRefreshed: lastRefreshed,
		Interval: Interval(m.Interval),
		Size: sizeFromString(m.OutputSize),
	}
}

type rawTimeSeriesIntraday struct {
	MetaData rawIntradayMetadata `json:"Meta Data"`
	TimeSeries1 rawTimeSeries `json:"Time Series (1min)"`
	TimeSeries5 rawTimeSeries `json:"Time Series (5min)"`
	TimeSeries15 rawTimeSeries `json:"Time Series (15min)"`
	TimeSeries30 rawTimeSeries `json:"Time Series (30min)"`
	TimeSeries60 rawTimeSeries `json:"Time Series (60min)"`
}

func (r *rawTimeSeriesIntraday) TimeSeries() rawTimeSeries {
	if len(r.TimeSeries1) > 0 {
		return r.TimeSeries1
	} else if len(r.TimeSeries5) > 0 {
		return r.TimeSeries5
	} else if len(r.TimeSeries15) > 0 {
		return r.TimeSeries15
	} else if len(r.TimeSeries30) > 0 {
		return r.TimeSeries30
	}
	return r.TimeSeries60
}


func (r *rawTimeSeriesIntraday) Parse() *TimeSeriesIntraday {
	metadata := r.MetaData.Parse()
	rawQuotes := r.TimeSeries()

	return &TimeSeriesIntraday{
		Metadata: metadata,
		Quotes: *rawQuotes.Parse(),
	}
}

type TimeSeriesIntraday struct {
	Metadata Metadata
	Quotes TimeSeries
}

func (a *AlphaVantage) TimeSeriesIntraday(symbol string, interval Interval, size Size) (*TimeSeriesIntraday, *ApiError) {
	params := map[string]string{
		"function": "TIME_SERIES_INTRADAY",
		"symbol": symbol,
		"interval": string(interval),
		"outputsize": string(size),
		"datatype": "json",
	}

	resp, apiError := a.request(params)
	if apiError != nil {
		return nil, apiError
	}

	raw := rawTimeSeriesIntraday{}
	err := json.Unmarshal(resp.Body, &raw)
	if err != nil {
		return nil, ToApiError(err, ERROR_RESPONSE_PARSE)
	}

	return raw.Parse(), nil
}
