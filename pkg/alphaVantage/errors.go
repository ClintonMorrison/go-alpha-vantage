package alphaVantage

import "fmt"

type ErrorType string

const (
	ERROR_REQUEST_FAILED ErrorType = "REQUEST_FAILED"
	ERROR_PARSE          ErrorType = "RESPONSE_PARSE"
	ERROR_RATE_LIMIT     ErrorType = "RATE_LIMIT"
	ERROR_OTHER          ErrorType = "OTHER"
)

type ApiError struct {
	Type ErrorType
	Message string
}

func (e ApiError) Error() string {
	return fmt.Sprintf("API Error - %s: %s", e.Type, e.Message)
}

func ToApiError(e error, errorType ErrorType) *ApiError {
	return &ApiError{
		Type: errorType,
		Message: e.Error()}
}