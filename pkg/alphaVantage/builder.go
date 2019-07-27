package alphaVantage

import (
	"net/http"
	"time"
)

type builder struct {
	key string
	baseUrl string
	httpClient *http.Client
	retryAttempts int // -1 for infinite
	retryBackoff time.Duration
}

func Builder() *builder {
	return &builder{
		key: "",
		baseUrl: "https://www.alphavantage.co/query",
		httpClient: &http.Client{},
		retryAttempts: 100,
		retryBackoff: 2 * time.Minute,
	}
}

func (b *builder) Key(key string) *builder {
	b.key = key
	return b
}

func (b *builder) BaseUrl(baseUrl string) *builder {
	b.baseUrl = baseUrl
	return b
}

func (b *builder) HttpClient(httpClient *http.Client) *builder {
	b.httpClient = httpClient
	return b
}

func (b *builder) RetryAttempts(retryAttempts int) *builder {
	b.retryAttempts = retryAttempts
	return b
}

func (b *builder) Build() *AlphaVantage {
	return &AlphaVantage{
		key: b.key,
		baseUrl: b.baseUrl,
		httpClient: b.httpClient,
		retryAttempts: b.retryAttempts,
		retryBackoff: b.retryBackoff,
	}
}