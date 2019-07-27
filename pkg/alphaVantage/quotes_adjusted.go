package alphaVantage

import (
	"github.com/ClintonMorrison/goAlphaVantage/internal/parse"
	"sort"
	"time"
)

type AdjustedQuote struct {
	Ticker           string
	Time             time.Time
	Open             float64
	High             float64
	Low              float64
	Close            float64
	AdjustedClose    float64
	Volume           float64
	DividendAmount   float64
	SplitCoefficient float64
}

type rawAdjustedDailyQuote struct {
	Open             string `json:"1. open"`
	High             string `json:"2. high"`
	Low              string `json:"3. low"`
	Close            string `json:"4. close"`
	AdjustedClose    string `json:"5. adjusted close"`
	Volume           string `json:"6. volume"`
	DividendAmount   string `json:"7. dividend amount"`
	SplitCoefficient string `json:"8. split coefficient"`
}

func (q *rawAdjustedDailyQuote) Parse(ticker string, t time.Time) *AdjustedQuote {
	return &AdjustedQuote{
		Ticker:           ticker,
		Time:             t,
		Open:             parse.FloatFromString(q.Open),
		High:             parse.FloatFromString(q.High),
		Low:              parse.FloatFromString(q.Low),
		AdjustedClose:    parse.FloatFromString(q.AdjustedClose),
		Close:            parse.FloatFromString(q.Close),
		Volume:           parse.FloatFromString(q.Volume),
		DividendAmount:   parse.FloatFromString(q.DividendAmount),
		SplitCoefficient: parse.FloatFromString(q.SplitCoefficient),
	}
}

type rawAdjustedQuote struct {
	Open           string `json:"1. open"`
	High           string `json:"2. high"`
	Low            string `json:"3. low"`
	Close          string `json:"4. close"`
	AdjustedClose  string `json:"5. adjusted close"`
	Volume         string `json:"6. volume"`
	DividendAmount string `json:"7. dividend amount"`
}

func (q *rawAdjustedQuote) Parse(ticker string, t time.Time) *AdjustedQuote {
	return &AdjustedQuote{
		Ticker:         ticker,
		Time:           t,
		Open:           parse.FloatFromString(q.Open),
		High:           parse.FloatFromString(q.High),
		Low:            parse.FloatFromString(q.Low),
		AdjustedClose:  parse.FloatFromString(q.AdjustedClose),
		Close:          parse.FloatFromString(q.Close),
		Volume:         parse.FloatFromString(q.Volume),
		DividendAmount: parse.FloatFromString(q.DividendAmount),
	}
}

type rawAdjustedDateSeries map[string]rawAdjustedDailyQuote
type rawAdjustedTimeSeries map[string]rawAdjustedQuote

type AdjustedTimeSeries map[time.Time]AdjustedQuote

func (ts *AdjustedTimeSeries) SortedTimes() []time.Time {
	var times timeSlice = make([]time.Time, 0, len(*ts))

	for t := range *ts {
		times = append(times, t)
	}

	sort.Sort(times)

	return times
}

func (ts AdjustedTimeSeries) Sorted() []AdjustedQuote {
	times := ts.SortedTimes()
	quotes := make([]AdjustedQuote, 0, len(times))

	for _, t := range times {
		quotes = append(quotes, ts[t])
	}

	return quotes
}

func (r *rawAdjustedDateSeries) Parse(ticker string) AdjustedTimeSeries {
	quotes := make(AdjustedTimeSeries, 0)

	for timeString, rawQuote := range *r {
		t := parse.DateFromString(timeString)
		quote := rawQuote.Parse(ticker, t)
		quotes[t] = *quote
	}

	return quotes
}

func (r *rawAdjustedTimeSeries) Parse(ticker string) AdjustedTimeSeries {
	quotes := make(AdjustedTimeSeries, 0)

	for timeString, rawQuote := range *r {
		t := parse.DateFromString(timeString)
		quote := rawQuote.Parse(ticker, t)
		quotes[t] = *quote
	}

	return quotes
}
