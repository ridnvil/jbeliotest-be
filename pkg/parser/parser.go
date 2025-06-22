package parser

import (
	"strconv"
	"strings"
	"time"
)

func ParseDate(input string) (time.Time, error) {
	t, err := time.Parse("01-02-06", input)
	if err != nil {
		return t, err
	}

	newDateFormat := t.Format("2006-01-02")
	newDate, err := time.Parse("2006-01-02", newDateFormat)
	return newDate, nil
}

func ParseStringInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func ParseDollarToFloat(s string) (float64, error) {
	clean := strings.ReplaceAll(s, "$", "")
	clean = strings.ReplaceAll(clean, ",", "")

	return strconv.ParseFloat(clean, 64)
}

func ParsePercentToFloat(s string) (float64, error) {
	clean := strings.TrimSuffix(s, "%")
	value, err := strconv.ParseFloat(clean, 64)
	if err != nil {
		return 0, err
	}
	return value / 100, nil
}
