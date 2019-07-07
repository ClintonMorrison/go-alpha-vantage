package alphaVantage

import (
	"encoding/json"
)

type rawTimeSeriesIntraday struct {
	TimeSeries1 rawTimeSeries `json:"Time Series (1min)"`
	TimeSeries5 rawTimeSeries `json:"Time Series (5min)"`
	TimeSeries15 rawTimeSeries `json:"Time Series (15min)"`
	TimeSeries30 rawTimeSeries `json:"Time Series (30min)"`
	TimeSeries60 rawTimeSeries `json:"Time Series (60min)"`
}

func (r *rawTimeSeriesIntraday) TimeSeries() rawTimeSeries {
	if len(r.TimeSeries1) > 0 {
		return r.TimeSeries1
	} else if len(r.TimeSeries5) > 0 {
		return r.TimeSeries5
	} else if len(r.TimeSeries15) > 0 {
		return r.TimeSeries15
	} else if len(r.TimeSeries30) > 0 {
		return r.TimeSeries30
	}
	return r.TimeSeries60
}


func (r rawTimeSeriesIntraday) Parse() TimeSeries {
	return r.TimeSeries().Parse()
}

func (a *AlphaVantage) TimeSeriesIntraday(symbol string, interval Interval, size Size) (TimeSeries, *ApiError) {
	params := map[string]string{
		"function": "TIME_SERIES_INTRADAY",
		"symbol": symbol,
		"interval": string(interval),
		"outputsize": string(size),
		"datatype": "json",
	}

	resp, apiError := a.request(params)
	if apiError != nil {
		return nil, apiError
	}

	raw := rawTimeSeriesIntraday{}
	err := json.Unmarshal(resp.Body, &raw)
	if err != nil {
		return nil, ToApiError(err, ERROR_PARSE)
	}

	return raw.Parse(), nil
}
