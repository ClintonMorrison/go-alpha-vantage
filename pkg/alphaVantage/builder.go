package alphaVantage

import "net/http"

type builder struct {
	key string
	baseUrl string
	httpClient *http.Client
}

func Builder() *builder {
	return &builder{
		key: "",
		baseUrl: "https://www.alphavantage.co/query",
		httpClient: &http.Client{},
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

func (b *builder) Build() *AlphaVantage {
	return &AlphaVantage{
		key: b.key,
		baseUrl: b.baseUrl,
		httpClient: b.httpClient,
	}
}