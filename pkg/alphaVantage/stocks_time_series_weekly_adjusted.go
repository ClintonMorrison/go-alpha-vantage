package alphaVantage

import (
"encoding/json"
)


type rawTimeSeriesWeeklyAdjusted struct {
	AdjustedTimeSeriesWeekly rawAdjustedTimeSeries `json:"Weekly Adjusted Time Series"`
}

func (r *rawTimeSeriesWeeklyAdjusted) Parse(ticker string) AdjustedTimeSeries {
	return r.AdjustedTimeSeriesWeekly.Parse(ticker)
}

func (a *AlphaVantage) TimeSeriesWeeklyAdjusted(symbol string) (AdjustedTimeSeries, *ApiError) {
	params := map[string]string{
		"function": "TIME_SERIES_WEEKLY_ADJUSTED",
		"symbol": symbol,
		"datatype": "json",
	}

	resp, apiError := a.request(params)
	if apiError != nil {
		return nil, apiError
	}

	raw := rawTimeSeriesWeeklyAdjusted{}
	err := json.Unmarshal(resp.Body, &raw)
	if err != nil {
		return nil, &ApiError{
			Type: ERROR_PARSE,
			Message: err.Error()}
	}

	return raw.Parse(symbol), nil
}
