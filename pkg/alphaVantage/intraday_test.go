package alphaVantage

import "testing"

func TestTimeSeriesIntraday_full(t *testing.T) {
	alphaVantage := clientForTest()

	resp, err := alphaVantage.TimeSeriesIntraday("MSFT", INTERVAL_30, SIZE_FULL)

	if err != nil {
		t.Errorf("API error: %s", err.Error())
		return
	}

	// Metadata
	assertStringEquals(t, "Intraday (30min) open, high, low, close prices and volume", resp.Metadata.Information)
	assertStringEquals(t, "MSFT", resp.Metadata.Symbol)
	assertStringEquals(t, "US/Eastern", resp.Metadata.LastRefreshed.Location().String())
	assertStringEquals(t, "30min", string(resp.Metadata.Interval))
	assertStringEquals(t, "full", string(resp.Metadata.Size))

	// Body
	date := *timeFromMap(resp.Quotes)
	assertNotZero(t, resp.Quotes[date].Open)
	assertNotZero(t, resp.Quotes[date].High)
	assertNotZero(t, resp.Quotes[date].Low)
	assertNotZero(t, resp.Quotes[date].Close)
	assertNotZero(t, resp.Quotes[date].Volume)
}

func TestTimeSeriesIntraday_compact(t *testing.T) {
	alphaVantage := clientForTest()

	resp, err := alphaVantage.TimeSeriesIntraday("MSFT", INTERVAL_1, SIZE_COMPACT)

	if err != nil {
		t.Errorf("API error: %s", err.Error())
		return
	}

	// Metadata
	assertStringEquals(t, "Intraday (1min) open, high, low, close prices and volume", resp.Metadata.Information)
	assertStringEquals(t, "MSFT", resp.Metadata.Symbol)
	assertStringEquals(t, "US/Eastern", resp.Metadata.LastRefreshed.Location().String())
	assertStringEquals(t, "1min", string(resp.Metadata.Interval))
	assertStringEquals(t, "compact", string(resp.Metadata.Size))

	// Body
	date := *timeFromMap(resp.Quotes)
	assertNotZero(t, resp.Quotes[date].Open)
	assertNotZero(t, resp.Quotes[date].High)
	assertNotZero(t, resp.Quotes[date].Low)
	assertNotZero(t, resp.Quotes[date].Close)
	assertNotZero(t, resp.Quotes[date].Volume)
}