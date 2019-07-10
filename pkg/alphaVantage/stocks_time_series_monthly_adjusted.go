package alphaVantage

import (
"encoding/json"
)


type rawTimeSeriesMonthlyAdjusted struct {
	AdjustedTimeSeriesMonthly rawAdjustedTimeSeries `json:"Monthly Adjusted Time Series"`
}

func (r *rawTimeSeriesMonthlyAdjusted) Parse() AdjustedTimeSeries {
	return r.AdjustedTimeSeriesMonthly.Parse()
}

func (a *AlphaVantage) TimeSeriesMonthlyAdjusted(symbol string) (AdjustedTimeSeries, *ApiError) {
	params := map[string]string{
		"function": "TIME_SERIES_MONTHLY_ADJUSTED",
		"symbol": symbol,
		"datatype": "json",
	}

	resp, apiError := a.request(params)
	if apiError != nil {
		return nil, apiError
	}

	raw := rawTimeSeriesMonthlyAdjusted{}
	err := json.Unmarshal(resp.Body, &raw)
	if err != nil {
		return nil, &ApiError{
			Type: ERROR_PARSE,
			Message: err.Error()}
	}

	return raw.Parse(), nil
}
