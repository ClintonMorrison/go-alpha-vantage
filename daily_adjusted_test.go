package goAlphaVantage

import (
	"testing"
)


func TestTimeSeriesDailyAdjusted_full(t *testing.T) {
	alphaVantage := clientForTest()

	resp, err := alphaVantage.TimeSeriesDailyAdjusted("MSFT", SIZE_FULL)

	if err != nil {
		t.Error(err)
	}

	// Metadata
	assertStringEquals(t, "Daily Time Series with Splits and Dividend Events", resp.Metadata.Information)
	assertStringEquals(t, "MSFT", resp.Metadata.Symbol)
	assertStringEquals(t, "US/Eastern", resp.Metadata.LastRefreshed.Location().String())
	assertStringEquals(t, "DAILY", string(resp.Metadata.Interval))
	assertStringEquals(t, "full", string(resp.Metadata.Size))

	// Body
	date := *timeFromAdjustedTimeSeries(resp.Quotes)
	assertNotZero(t, resp.Quotes[date].Open)
	assertNotZero(t, resp.Quotes[date].High)
	assertNotZero(t, resp.Quotes[date].Low)
	assertNotZero(t, resp.Quotes[date].Close)
	assertNotZero(t, resp.Quotes[date].AdjustedClose)
	assertNotZero(t, resp.Quotes[date].Volume)
	assertNotZero(t, resp.Quotes[date].SplitCoefficient)
}

func TestTimeSeriesDailyAdjusted_compact(t *testing.T) {
	alphaVantage := clientForTest()

	resp, err := alphaVantage.TimeSeriesDailyAdjusted("MSFT", SIZE_COMPACT)

	if err != nil {
		t.Error(err)
	}

	// Metadata
	assertStringEquals(t, "Daily Time Series with Splits and Dividend Events", resp.Metadata.Information)
	assertStringEquals(t, "MSFT", resp.Metadata.Symbol)
	assertStringEquals(t, "US/Eastern", resp.Metadata.LastRefreshed.Location().String())
	assertStringEquals(t, "DAILY", string(resp.Metadata.Interval))
	assertStringEquals(t, "compact", string(resp.Metadata.Size))

	// Body
	date := *timeFromAdjustedTimeSeries(resp.Quotes)
	assertNotZero(t, resp.Quotes[date].Open)
	assertNotZero(t, resp.Quotes[date].High)
	assertNotZero(t, resp.Quotes[date].Low)
	assertNotZero(t, resp.Quotes[date].Close)
	assertNotZero(t, resp.Quotes[date].AdjustedClose)
	assertNotZero(t, resp.Quotes[date].Volume)
	assertNotZero(t, resp.Quotes[date].SplitCoefficient)
}

