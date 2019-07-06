package goAlphaVantage

import (
	"net/http"
	"fmt"
	"strings"
	"io/ioutil"
)

type Interval string
const (
	INTERVAL_1 Interval = "1min"
	INTERVAL_5 Interval = "5min"
	INTERVAL_15 Interval = "5min"
	INTERVAL_30 Interval = "30min"
	INTERVAL_60 Interval = "360min"
)

func intervalFromString(s string) Interval {
	switch s {
	case "1min": return INTERVAL_1
	case "5min": return INTERVAL_5
	case "15min": return INTERVAL_15
	case "30min": return INTERVAL_30
	case "60min": return INTERVAL_60
	default: return INTERVAL_60
	}
}

type Size string
const (
	SIZE_COMPACT Size = "compact"
	SIZE_FULL Size = "full"
)

func sizeFromString(s string) Size {
	switch s {
	case "Full size": return SIZE_FULL
	default: return SIZE_COMPACT

	}
}

type AlphaVantage struct {
	key string
	baseUrl string
	httpClient *http.Client
}

type rawResponse struct {
	Code int
	Body []byte
}

func (a *AlphaVantage) request(params map[string]string) (*rawResponse, error) {
	params["apikey"] = a.key
	params["datatype"] = "json"

	query := toQuery(params)
	url := fmt.Sprintf("%s?%s", a.baseUrl, query)

	fmt.Println("Request: " + url)
	resp, err := a.httpClient.Get(url)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &rawResponse{
		Code: resp.StatusCode,
		Body: body,
	}, nil
}


func toQuery(params map[string]string) string {
	if len(params) == 0 {
		return ""
	}

	fields := make([]string, 0, len(params))
	for name, value := range params {
		fields = append(fields, fmt.Sprintf("%s=%s", name, value))
	}

	return strings.Join(fields, "&")
}

