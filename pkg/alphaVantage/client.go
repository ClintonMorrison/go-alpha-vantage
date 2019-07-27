package alphaVantage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type AlphaVantageClient struct {
	key           string
	baseUrl       string
	httpClient    *http.Client
	retryAttempts int // -1 for infinite
	retryBackoff  time.Duration
}

func Client() *AlphaVantageClient {
	return &AlphaVantageClient{
		key:           "",
		baseUrl:       "https://www.alphavantage.co/query",
		httpClient:    &http.Client{},
		retryAttempts: 100,
		retryBackoff:  2 * time.Minute,
	}
}

func (c *AlphaVantageClient) Key(key string) *AlphaVantageClient {
	c.key = key
	return c
}

func (c *AlphaVantageClient) BaseUrl(baseUrl string) *AlphaVantageClient {
	c.baseUrl = baseUrl
	return c
}

func (c *AlphaVantageClient) HttpClient(httpClient *http.Client) *AlphaVantageClient {
	c.httpClient = httpClient
	return c
}

func (c *AlphaVantageClient) RetryAttempts(retryAttempts int) *AlphaVantageClient {
	c.retryAttempts = retryAttempts
	return c
}

type rawResponse struct {
	Code int
	Body []byte
}

type rawErrorResponse struct {
	Note string `json:"Note"`
}

func (e rawErrorResponse) Error() string {
	return e.Note
}

func (e rawErrorResponse) ToApiError() *ApiError {
	errorType := ERROR_OTHER

	if strings.Contains(e.Note, "API call frequency") {
		errorType = ERROR_RATE_LIMIT
	}

	return &ApiError{
		Type:    errorType,
		Message: e.Note}
}

func (a *AlphaVantageClient) request(params map[string]string) (*rawResponse, *ApiError) {
	var resp *rawResponse
	var apiError *ApiError

	retries := 0
	for retries <= a.retryAttempts {
		resp, apiError = a.internalRequest(params)

		if apiError != nil && apiError.Type == ERROR_RATE_LIMIT {
			fmt.Println("Retries", retries)
			retries += 1
			time.Sleep(a.retryBackoff)
			continue
		}

		break
	}

	return resp, apiError
}

func (a *AlphaVantageClient) internalRequest(params map[string]string) (*rawResponse, *ApiError) {
	params["apikey"] = a.key
	params["datatype"] = "json"

	params["symbol"] = strings.Replace(params["symbol"], ".", "-", -1) // "TSE:ATD.B" -> "TSE:ATD-B"

	query := toQuery(params)
	url := fmt.Sprintf("%s?%s", a.baseUrl, query)

	fmt.Println("Request: " + url)
	resp, err := a.httpClient.Get(url)

	if err != nil {
		return nil, &ApiError{
			Type:    ERROR_REQUEST_FAILED,
			Message: err.Error()}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, ToApiError(err, ERROR_PARSE)
	}

	// Check for error
	errResponse := &rawErrorResponse{}
	json.Unmarshal(body, errResponse)
	if len(errResponse.Error()) != 0 {
		return nil, errResponse.ToApiError()
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
		fields = append(fields, fmt.Sprintf("%s=%s", name, url.QueryEscape(value)))
	}

	return strings.Join(fields, "&")
}
