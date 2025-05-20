package utils

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func ParseDuration(input string) (time.Duration, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return 0, errors.New("empty duration string")
	}

	if strings.HasSuffix(input, "d") {
		dayStr := strings.TrimSuffix(input, "d")
		days, err := strconv.Atoi(dayStr)
		if err != nil {
			return 0, errors.New("invalid day duration format")
		}
		return time.Duration(days) * 24 * time.Hour, nil
	}

	// Fallback to time.ParseDuration for valid time strings
	return time.ParseDuration(input)
}
