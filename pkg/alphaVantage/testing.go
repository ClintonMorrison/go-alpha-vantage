package alphaVantage

import (
	"testing"
	"time"
)

func assertStringEquals(t *testing.T, expected string, actual string) {
	if expected != actual {
		t.Errorf("expected '%s' but was '%s'", expected, actual)
	}
}

func assertFloatEquals(t *testing.T, expected float64, actual float64) {
	if expected != actual {
		t.Errorf("expected '%f' but was '%f'", expected, actual)
	}
}

func assertNotZero(t *testing.T, actual float64) {
	if actual == 0 {
		t.Errorf("expected '%f' to not be 0", actual)
	}
}

func assertGreaterThan(t *testing.T, min int, actual int) {
	if actual < min {
		t.Errorf("expected '%d' to be greater than %d", actual, min)
	}
}

func clientForTest() *AlphaVantageClient {
	return Client().Key("TODO-KEY-HERE")
}

func timeFromMap(series TimeSeries) *time.Time {
	for t, _ := range series {
		return &t
	}
	t := time.Now()
	return &t
}

func timeFromAdjustedTimeSeries(series AdjustedTimeSeries) *time.Time {
	for t, _ := range series {
		return &t
	}
	t := time.Now()
	return &t
}
