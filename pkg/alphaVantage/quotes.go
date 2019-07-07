package alphaVantage

import (
	"time"
	"github.com/ClintonMorrison/goAlphaVantage/internal/parse"
)

type Quote struct {
	Open float64
	High float64
	Low float64
	Close float64
	Volume float64
}

type AdjustedQuote struct {
	Open float64
	High float64
	Low float64
	Close float64
	AdjustedClose float64
	Volume float64
	DividendAmount float64
	SplitCoefficient float64
}


type rawQuote struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

func (q *rawQuote) Parse() *Quote {
	return &Quote{
		Open: parse.FloatFromString(q.Open),
		High: parse.FloatFromString(q.High),
		Low: parse.FloatFromString(q.Low),
		Close: parse.FloatFromString(q.Close),
		Volume: parse.FloatFromString(q.Volume),
	}
}

type rawAdjustedQuote struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	AdjustedClose  string `json:"5. adjusted close"`
	Volume string `json:"6. volume"`
	DividendAmount string `json:"7. dividend amount"`
	SplitCoefficient string `json:"8. split coefficient"`
}

func (q *rawAdjustedQuote) Parse() *AdjustedQuote {
	return &AdjustedQuote{
		Open: parse.FloatFromString(q.Open),
		High: parse.FloatFromString(q.High),
		Low: parse.FloatFromString(q.Low),
		AdjustedClose: parse.FloatFromString(q.AdjustedClose),
		Close: parse.FloatFromString(q.Close),
		Volume: parse.FloatFromString(q.Volume),
		DividendAmount: parse.FloatFromString(q.DividendAmount),
		SplitCoefficient: parse.FloatFromString(q.SplitCoefficient),
	}
}

type rawTimeSeries map[string]rawQuote

type TimeSeries map[time.Time]Quote

func (r rawTimeSeries) Parse() TimeSeries {
	quotes := make(TimeSeries, 0)

	for timeString, rawQuote := range r {
		quote := rawQuote.Parse()
		t, _ := parse.TimeFromString(timeString)
		quotes[t] = *quote
	}

	return quotes
}


type rawAdjustedTimeSeries map[string]rawAdjustedQuote

type AdjustedTimeSeries map[time.Time]AdjustedQuote

func (r *rawAdjustedTimeSeries) Parse() AdjustedTimeSeries {
	quotes := make(AdjustedTimeSeries, 0)

	for timeString, rawQuote := range *r {
		quote := rawQuote.Parse()
		t, _ := parse.TimeFromString(timeString)
		quotes[t] = *quote
	}

	return quotes
}


