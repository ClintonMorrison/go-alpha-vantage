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


	quote, err := client.Quote("TSE:SHOP")

	if err != nil {
		panic(err)
	}

	fmt.Println(quote)

}