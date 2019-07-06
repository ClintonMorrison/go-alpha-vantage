package main

import (
	"net/http"
	"net"
	"time"
	"github.com/ClintonMorrison/goAlphaVantage"
	"github.com/ClintonMorrison/goAlphaVantage/config"
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

	alphaVantage := goAlphaVantage.Builder().
		Key(config.ALPHA_VANTAGE_KEY).
		HttpClient(httpClient).
		Build()


	_, err := alphaVantage.TimeSeriesIntraday(
		"MSFT",
		goAlphaVantage.INTERVAL_30,
		goAlphaVantage.SIZE_FULL)

	if err != nil {
		panic(err)
	}
}