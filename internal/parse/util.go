package parse

import (
	"time"
	"strconv"
	"fmt"
	"strings"
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

func DateFromString(t string) (time.Time) {
	result, err :=  time.Parse("2006-01-02", t)

	if err != nil {
		fmt.Printf("Error parsing date: '%s': %s", t, err.Error())
	}

	return result
}


func FloatFromString(s string) float64 {
	val, err := strconv.ParseFloat(s, 64)

	if err != nil {
		fmt.Printf("Error parsing float '%s': %s\n", s, err.Error())
		return 0
	}

	return val
}

func FloatFromPercentString(s string) float64 {
	cleaned := strings.Trim(s, "%")
	return FloatFromString(cleaned) / 100
}