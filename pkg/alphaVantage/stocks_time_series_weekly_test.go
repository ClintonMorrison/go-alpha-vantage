package alphaVantage

import (
	"testing"
)

func TestStocksTimeSeriesWeekly(t *testing.T) {
	alphaVantage := clientForTest()

	quotes, err := alphaVantage.TimeSeriesWeekly("MSFT")

	if err != nil {
		t.Errorf("API error: %s", err.Error())
		return
	}

	// Make sure multiple quotes returned
	assertGreaterThan(t, 10, len(quotes))

	// Body
	date := *timeFromMap(quotes)
	assertStringEquals(t, "MSFT", quotes[date].Ticker)
	assertNotZero(t, float64(quotes[date].Time.Unix()))
	assertNotZero(t, quotes[date].Open)
	assertNotZero(t, quotes[date].High)
	assertNotZero(t, quotes[date].Low)
	assertNotZero(t, quotes[date].Close)
	assertNotZero(t, quotes[date].Volume)
}
