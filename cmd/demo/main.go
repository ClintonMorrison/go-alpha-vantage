package main

import (
	"net/http"
	"net"
	"time"
	"github.com/ClintonMorrison/goAlphaVantage/pkg/alphaVantage"
	"github.com/ClintonMorrison/goAlphaVantage/config"
	"fmt"
)


func main() {
	var httpTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	var httpClient = &http.Client{
		Timeout: time.Second * 10,
		Transport: httpTransport,
	}

	client := alphaVantage.Builder().
		Key(config.ALPHA_VANTAGE_KEY).
		HttpClient(httpClient).
		Build()


	quotes, err := client.TimeSeriesIntraday(
		"TSE:SHOP",
		alphaVantage.INTERVAL_30,
		alphaVantage.SIZE_FULL)

	if err != nil {
		panic(err)
	}

	for date, quote := range quotes {
		fmt.Printf("%s - %f\n", date.String(), quote.Close)
	}
}