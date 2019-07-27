package alphaVantage

import (
	"encoding/json"
	"github.com/ClintonMorrison/goAlphaVantage/internal/parse"
	"time"
)

type rawGlobalQuote struct {
	Symbol           string `json:"01. symbol"`
	Open             string `json:"02. open"`
	High             string `json:"03. high"`
	Low              string `json:"04. low"`
	Price            string `json:"05. price"`
	Volume           string `json:"06. volume"`
	LatestTradingDay string `json:"07. latest trading day"`
	PreviousClose    string `json:"08. previous close"`
	ChangePercent    string `json:"10. change percent"`
}

type RealTimeQuote struct {
	Ticker            string
	Time              time.Time
	Open              float64
	High              float64
	Low               float64
	Current           float64
	Volume            float64
	LatestTradingDate time.Time
	PreviousClose     float64
	PercentChange     float64
}

func (q *rawGlobalQuote) Parse(ticker string, t time.Time) RealTimeQuote {
	return RealTimeQuote{
		Ticker:            ticker,
		Time:              t,
		Open:              parse.FloatFromString(q.Open),
		High:              parse.FloatFromString(q.High),
		Low:               parse.FloatFromString(q.Low),
		Current:           parse.FloatFromString(q.Price),
		Volume:            parse.FloatFromString(q.Volume),
		LatestTradingDate: parse.DateFromString(q.LatestTradingDay),
		PreviousClose:     parse.FloatFromString(q.PreviousClose),
		PercentChange:     parse.FloatFromPercentString(q.ChangePercent)}
}

type rawQuoteResponse struct {
	Quote rawGlobalQuote `json:"Global Quote"`
}

func (r *rawQuoteResponse) Parse(ticker string, t time.Time) *RealTimeQuote {
	quote := r.Quote.Parse(ticker, t)
	return &quote
}

func (a *AlphaVantageClient) Quote(symbol string) (*RealTimeQuote, *ApiError) {
	params := map[string]string{
		"function": "GLOBAL_QUOTE",
		"symbol":   symbol,
		"datatype": "json",
	}

	resp, apiError := a.request(params)
	if apiError != nil {
		return nil, apiError
	}

	raw := rawQuoteResponse{}
	err := json.Unmarshal(resp.Body, &raw)
	if err != nil {
		return nil, ToApiError(err, ERROR_PARSE)
	}

	return raw.Parse(symbol, time.Now()), nil
}
