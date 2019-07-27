# goAlphaVantage
Golang API client for [AlphaVantage](https://www.alphavantage.co/documentation/).

---

To need a key to use the API. You can get one for free [here](https://www.alphavantage.co/support/#api-key).

To install this library:
```
go get github.com/ClintonMorrison/goAlphaVantage
```

Instantiate a client:
```
client := alphaVantage.Client().
    Key("your-alpha-vantage-key")
```

Most method calls and the data they return match the documented AlphaVantage [endpoints](https://www.alphavantage.co/documentation/).

#### Examples

Getting a real time quote for Microsoft:
```
quote, err := client.Quote("MSFT")

// Prints "MSFT: 141.340000"
fmt.Printf("%s: %f", quote.Ticker, quote.Current) 
```

Printing monthly historical prices for Shopify (on the TSX):
```
quotes, err := client.TimeSeriesMonthlyAdjusted("TSE:SHOP")

for _, quote := range quotes.Sorted() {
    fmt.Printf("%s: %f\n", quote.Time.Format("2006-01-02"), quote.Close)
}
```
