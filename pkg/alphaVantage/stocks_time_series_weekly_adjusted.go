package alphaVantage

import (
	"encoding/json"
)

type rawTimeSeriesWeeklyAdjusted struct {
	AdjustedDateSeries rawAdjustedDateSeries `json:"Weekly Adjusted Time Series"`
}

func (r *rawTimeSeriesWeeklyAdjusted) Parse(ticker string) AdjustedTimeSeries {
	return r.AdjustedDateSeries.Parse(ticker)
}

func (a *AlphaVantageClient) TimeSeriesWeeklyAdjusted(symbol string) (AdjustedTimeSeries, *ApiError) {
	params := map[string]string{
		"function": "TIME_SERIES_WEEKLY_ADJUSTED",
		"symbol":   symbol,
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
			Type:    ERROR_PARSE,
			Message: err.Error()}
	}

	return raw.Parse(symbol), nil
}
