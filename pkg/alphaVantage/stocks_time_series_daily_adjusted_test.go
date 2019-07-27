package alphaVantage

import (
	"testing"
)


func TestStocksTimeSeriesDailyAdjusted_full(t *testing.T) {
	alphaVantage := clientForTest()

	quotes, err := alphaVantage.TimeSeriesDailyAdjusted("MSFT", SIZE_FULL)

	if err != nil {
		t.Errorf("API error: %s", err.Error())
		return
	}

	// Body
	date := *timeFromAdjustedTimeSeries(quotes)
	assertStringEquals(t, "MSFT", quotes[date].Ticker)
	assertNotZero(t, float64(quotes[date].Time.Unix()))
	assertNotZero(t, quotes[date].Open)
	assertNotZero(t, quotes[date].High)
	assertNotZero(t, quotes[date].Low)
	assertNotZero(t, quotes[date].Close)
	assertNotZero(t, quotes[date].AdjustedClose)
	assertNotZero(t, quotes[date].Volume)
	assertNotZero(t, quotes[date].SplitCoefficient)
}

func TestStocksTimeSeriesDailyAdjusted_compact(t *testing.T) {
	alphaVantage := clientForTest()

	quotes, err := alphaVantage.TimeSeriesDailyAdjusted("MSFT", SIZE_COMPACT)

	if err != nil {
		t.Errorf("API error: %s", err.Error())
		return
	}

	// Body
	date := *timeFromAdjustedTimeSeries(quotes)
	assertNotZero(t, quotes[date].Open)
	assertNotZero(t, quotes[date].High)
	assertNotZero(t, quotes[date].Low)
	assertNotZero(t, quotes[date].Close)
	assertNotZero(t, quotes[date].AdjustedClose)
	assertNotZero(t, quotes[date].Volume)
	assertNotZero(t, quotes[date].SplitCoefficient)
}

