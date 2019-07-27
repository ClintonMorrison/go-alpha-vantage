package alphaVantage

import (
"encoding/json"
	"fmt"
)


type rawTimeSeriesMonthlyAdjusted struct {
	AdjustedTimeSeriesMonthly rawAdjustedTimeSeries `json:"Monthly Adjusted Time Series"`
}

func (r *rawTimeSeriesMonthlyAdjusted) Parse(ticker string) AdjustedTimeSeries {
	return r.AdjustedTimeSeriesMonthly.Parse(ticker)
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

	fmt.Println("Returning", raw.Parse(symbol), "|", nil)

	return raw.Parse(symbol), nil
}
