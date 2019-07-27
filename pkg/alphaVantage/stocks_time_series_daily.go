package alphaVantage

import (
	"encoding/json"
)

type rawTimeSeriesDaily struct {
	DateSeries rawDateSeries `json:"Time Series (Daily)"`
}

func (r *rawTimeSeriesDaily) Parse(ticker string) TimeSeries {
	return r.DateSeries.Parse(ticker)
}

func (a *AlphaVantageClient) TimeSeriesDaily(symbol string, size Size) (TimeSeries, *ApiError) {
	params := map[string]string{
		"function":   "TIME_SERIES_DAILY",
		"symbol":     symbol,
		"outputsize": string(size),
		"datatype":   "json",
	}

	resp, apiError := a.request(params)
	if apiError != nil {
		return nil, apiError
	}

	raw := rawTimeSeriesDaily{}
	err := json.Unmarshal(resp.Body, &raw)
	if err != nil {
		return nil, ToApiError(err, ERROR_PARSE)
	}

	return raw.Parse(symbol), nil
}
