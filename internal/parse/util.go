package parse

import (
	"time"
	"strconv"
	"fmt"
)

func TimeFromString(t string) (time.Time, error) {
	return time.Parse("2006-01-02 15:04:05", t)
}

func TimeFromStringLocation(t string, location *time.Location) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", t, location)
}

func DateFromStringLocation(t string, location *time.Location) (time.Time, error) {
	return time.ParseInLocation("2006-01-02", t, location)
}

func FloatFromString(s string) float64 {
	val, err := strconv.ParseFloat(s, 64)

	if err != nil {
		fmt.Printf("Error parsing float '%s': %s\n", s, err.Error())
		return 0
	}

	return val
}