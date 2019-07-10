package alphaVantage

import (
"encoding/json"
)


type rawTimeSeriesDailyAdjusted struct {
	AdjustedTimeSeriesDaily rawAdjustedDailyTimeSeries `json:"Time Series (Daily)"`
}

func (r *rawTimeSeriesDailyAdjusted) Parse() AdjustedTimeSeries {
	return r.AdjustedTimeSeriesDaily.Parse()
}

func (a *AlphaVantage) TimeSeriesDailyAdjusted(symbol string, size Size) (AdjustedTimeSeries, *ApiError) {
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
			Type: ERROR_PARSE,
			Message: err.Error()}
	}

	return raw.Parse(), nil
}
