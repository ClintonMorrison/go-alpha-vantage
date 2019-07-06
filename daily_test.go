package goAlphaVantage

import (
	"testing"
)

func TestTimeSeriesDaily_full(t *testing.T) {
	alphaVantage := clientForTest()

	resp, err := alphaVantage.TimeSeriesDaily("MSFT", SIZE_FULL)

	if err != nil {
		t.Error(err)
	}

	// Metadata
	assertStringEquals(t, "Daily Prices (open, high, low, close) and Volumes", resp.Metadata.Information)
	assertStringEquals(t, "MSFT", resp.Metadata.Symbol)
	assertStringEquals(t, "US/Eastern", resp.Metadata.LastRefreshed.Location().String())
	assertStringEquals(t, "DAILY", string(resp.Metadata.Interval))
	assertStringEquals(t, "full", string(resp.Metadata.Size))

	// Body
	date := *timeFromMap(resp.Quotes)
	assertNotZero(t, resp.Quotes[date].Open)
	assertNotZero(t, resp.Quotes[date].High)
	assertNotZero(t, resp.Quotes[date].Low)
	assertNotZero(t, resp.Quotes[date].Close)
	assertNotZero(t, resp.Quotes[date].Volume)
}

func TestTimeSeriesDaily_compact(t *testing.T) {
	alphaVantage := clientForTest()

	resp, err := alphaVantage.TimeSeriesDaily("MSFT", SIZE_COMPACT)

	if err != nil {
		t.Error(err)
	}

	// Metadata
	assertStringEquals(t, "Daily Prices (open, high, low, close) and Volumes", resp.Metadata.Information)
	assertStringEquals(t, "MSFT", resp.Metadata.Symbol)
	assertStringEquals(t, "US/Eastern", resp.Metadata.LastRefreshed.Location().String())
	assertStringEquals(t, "DAILY", string(resp.Metadata.Interval))
	assertStringEquals(t, "compact", string(resp.Metadata.Size))

	// Body
	date := *timeFromMap(resp.Quotes)
	assertNotZero(t, resp.Quotes[date].Open)
	assertNotZero(t, resp.Quotes[date].High)
	assertNotZero(t, resp.Quotes[date].Low)
	assertNotZero(t, resp.Quotes[date].Close)
	assertNotZero(t, resp.Quotes[date].Volume)
}

