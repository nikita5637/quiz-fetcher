package sixty_seconds

import (
	"time"

	time_utils "github.com/nikita5637/quiz-fetcher/utils/time"
)

const (
	timeFormatString = "2 Jan 2006, 15:04 p.m."
)

func convertDateTime(dateTime string) (time.Time, error) {
	loc, err := time.LoadLocation(time_utils.TimeZoneMoscow)
	if err != nil {
		return time.Time{}, nil
	}

	t, err := time.ParseInLocation(timeFormatString, dateTime, loc)
	if err != nil {
		return time.Time{}, err
	}

	return t.UTC().Add(12 * time.Hour), nil
}
