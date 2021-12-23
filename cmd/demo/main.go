package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/ClintonMorrison/goAlphaVantage/pkg/alphaVantage"
)

func main() {
	var httpTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	var httpClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: httpTransport,
	}

	client := alphaVantage.Client().
		Key("ADD-KEY-HERE").
		HttpClient(httpClient)

	quotes, err := client.TimeSeriesMonthly("TSE:SHOP")

	if err != nil {
		panic(err)
	}

	for _, quote := range quotes.Sorted() {
		fmt.Printf("%s: %f\n", quote.Time.Format("2006-01-02"), quote.Close)
	}
}
