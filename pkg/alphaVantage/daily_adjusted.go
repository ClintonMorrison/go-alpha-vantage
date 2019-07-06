package alphaVantage

import (
"encoding/json"
)


type rawTimeSeriesDailyAdjusted struct {
	MetaData rawDailyMetadata `json:"Meta Data"`
	AdjustedTimeSeriesDaily rawAdjustedTimeSeries `json:"Time Series (Daily)"`
}


func (r *rawTimeSeriesDailyAdjusted) Parse() *TimeSeriesDailyAdjusted {
	metadata := r.MetaData.Parse()
	rawQuotes := r.AdjustedTimeSeriesDaily

	return &TimeSeriesDailyAdjusted{
		Metadata: metadata,
		Quotes: *rawQuotes.Parse(),
	}
}

type TimeSeriesDailyAdjusted struct {
	Metadata Metadata
	Quotes AdjustedTimeSeries
}

func (a *AlphaVantage) TimeSeriesDailyAdjusted(symbol string, size Size) (*TimeSeriesDailyAdjusted, error) {
	params := map[string]string{
		"function": "TIME_SERIES_DAILY_ADJUSTED",
		"symbol": symbol,
		"outputsize": string(size),
		"datatype": "json",
	}

	resp, err := a.request(params)
	if err != nil {
		return nil, err
	}


	raw := rawTimeSeriesDailyAdjusted{}
	err = json.Unmarshal(resp.Body, &raw)
	if err != nil {
		return nil, err
	}

	return raw.Parse(), err
}
