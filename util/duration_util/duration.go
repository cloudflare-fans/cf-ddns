package duration_util

import (
	"errors"
	"regexp"
	"strconv"
	"time"
)

func ParseNumTypedDuration(s string) (time.Duration, error) {
	re := regexp.MustCompile(`^(\d+)([smhdwy])$`)
	matches := re.FindStringSubmatch(s)
	if len(matches) != 3 {
		return 0, errors.New("invalid duration format")
	}

	num, err := strconv.Atoi(matches[1])

	if err != nil {
		return 0, err
	}

	var duration time.Duration

	switch matches[2] {
	case "s":
		duration = time.Duration(num) * time.Second
	case "m":
		duration = time.Duration(num) * time.Minute
	case "h":
		duration = time.Duration(num) * time.Hour
	case "d":
		duration = time.Duration(num) * 24 * time.Hour
	case "w":
		duration = time.Duration(num) * 7 * 24 * time.Hour
	case "y":
		duration = time.Duration(num) * 365 * 24 * time.Hour
	default:
		return 0, errors.New("unsupported time unit")
	}
	return duration, nil
}
