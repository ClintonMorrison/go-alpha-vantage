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

func (a *AlphaVantage) TimeSeriesDailyAdjusted(symbol string, size Size) (*TimeSeriesDailyAdjusted, *ApiError) {
	params := map[string]string{
		"function": "TIME_SERIES_DAILY_ADJUSTED",
		"symbol": symbol,
		"outputsize": string(size),
		"datatype": "json",
	}

	resp, apiError := a.request(params)
	if apiError != nil {
		return nil, apiError
	}


	raw := rawTimeSeriesDailyAdjusted{}
	err := json.Unmarshal(resp.Body, &raw)
	if err != nil {
		return nil, &ApiError{
			Type: ERROR_RESPONSE_PARSE,
			Message: err.Error()}
	}

	return raw.Parse(), nil
}
