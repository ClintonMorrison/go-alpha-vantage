package alphaVantage

import "testing"

func TestStocksQuote(t *testing.T) {
	alphaVantage := clientForTest()

	quote, err := alphaVantage.Quote("MSFT")

	if err != nil {
		t.Errorf("API error: %s", err.Error())
		return
	}

	assertStringEquals(t, "MSFT", quote.Ticker)
	assertNotZero(t, float64(quote.Time.Unix()))
	assertNotZero(t, quote.Open)
	assertNotZero(t, quote.High)
	assertNotZero(t, quote.Low)
	assertNotZero(t, quote.Current)
	assertNotZero(t, quote.Volume)
	assertNotZero(t, float64(quote.LatestTradingDate.Unix()))
	assertNotZero(t, quote.PreviousClose)
	assertNotZero(t, quote.PercentChange)
}
