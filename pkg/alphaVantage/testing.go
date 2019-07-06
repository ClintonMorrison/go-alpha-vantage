package alphaVantage

import (
	"testing"
	"time"
	"github.com/ClintonMorrison/goAlphaVantage/internal/config"
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

func clientForTest() *AlphaVantage {
	return Builder().Key(config.ALPHA_VANTAGE_KEY).Build()
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

