package alphaVantage

import (
	"net/http"
	"fmt"
	"strings"
	"io/ioutil"
	"encoding/json"
)

type AlphaVantage struct {
	key string
	baseUrl string
	httpClient *http.Client
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

	// Check for error
	errResponse := &rawErrorResponse{}
	json.Unmarshal(body, errResponse)
	if len(errResponse.Error()) != 0 {
		return nil, errResponse
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

