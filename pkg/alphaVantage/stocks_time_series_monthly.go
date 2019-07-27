package alphaVantage

import (
	"encoding/json"
)

type rawTimeSeriesMonthly struct {
	DateSeries rawDateSeries `json:"Monthly Time Series"`
}

func (r *rawTimeSeriesMonthly) Parse(ticker string) TimeSeries {
	return r.DateSeries.Parse(ticker)
}

func (a *AlphaVantageClient) TimeSeriesMonthly(symbol string) (TimeSeries, *ApiError) {
	params := map[string]string{
		"function": "TIME_SERIES_MONTHLY",
		"symbol":   symbol,
		"datatype": "json",
	}

	resp, apiError := a.request(params)
	if apiError != nil {
		return nil, apiError
	}

	raw := rawTimeSeriesMonthly{}
	err := json.Unmarshal(resp.Body, &raw)
	if err != nil {
		return nil, ToApiError(err, ERROR_PARSE)
	}

	return raw.Parse(symbol), nil
}
