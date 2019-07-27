package alphaVantage

import "testing"

func TestStocksTimeSeriesIntraday_full(t *testing.T) {
	alphaVantage := clientForTest()

	quotes, err := alphaVantage.TimeSeriesIntraday("MSFT", INTERVAL_30, SIZE_FULL)

	if err != nil {
		t.Errorf("API error: %s", err.Error())
		return
	}

	date := *timeFromMap(quotes)
	assertStringEquals(t, "MSFT", quotes[date].Ticker)
	assertNotZero(t, float64(quotes[date].Time.Unix()))
	assertNotZero(t, quotes[date].Open)
	assertNotZero(t, quotes[date].High)
	assertNotZero(t, quotes[date].Low)
	assertNotZero(t, quotes[date].Close)
	assertNotZero(t, quotes[date].Volume)
}

func TestStocksTimeSeriesIntraday_compact(t *testing.T) {
	alphaVantage := clientForTest()

	quotes, err := alphaVantage.TimeSeriesIntraday("MSFT", INTERVAL_1, SIZE_COMPACT)

	if err != nil {
		t.Errorf("API error: %s", err.Error())
		return
	}

	// Body
	date := *timeFromMap(quotes)
	assertNotZero(t, quotes[date].Open)
	assertNotZero(t, quotes[date].High)
	assertNotZero(t, quotes[date].Low)
	assertNotZero(t, quotes[date].Close)
	assertNotZero(t, quotes[date].Volume)
}