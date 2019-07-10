package alphaVantage

import (
"encoding/json"
)

type rawTimeSeriesWeekly struct {
	TimeSeriesDaily rawTimeSeries `json:"Weekly Time Series"`
}

func (r *rawTimeSeriesWeekly) Parse() TimeSeries {
	return r.TimeSeriesDaily.Parse()
}

func (a *AlphaVantage) TimeSeriesWeekly(symbol string) (TimeSeries, *ApiError) {
	params := map[string]string{
		"function": "TIME_SERIES_WEEKLY",
		"symbol": symbol,
		"datatype": "json",
	}

	resp, apiError := a.request(params)
	if apiError != nil {
		return nil, apiError
	}

	raw := rawTimeSeriesWeekly{}
	err := json.Unmarshal(resp.Body, &raw)
	if err != nil {
		return nil, ToApiError(err, ERROR_PARSE)
	}

	return raw.Parse(), nil
}
