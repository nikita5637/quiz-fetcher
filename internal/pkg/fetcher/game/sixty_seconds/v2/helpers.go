package sixty_seconds

import (
	"strings"
	"time"

	time_utils "github.com/nikita5637/quiz-fetcher/utils/time"
)

const (
	timeFormatStringEng = "2 Jan 2006, 15:04 p.m."
	timeFormatStringRus = "2 Jan 2006 15:04"
)

var (
	months = map[string]string{
		"января":   "Jan",
		"февраля":  "Feb",
		"марта":    "Mar",
		"апреля":   "Apr",
		"мая":      "May",
		"июня":     "Jun",
		"июля":     "Jul",
		"августа":  "Aug",
		"сентября": "Sep",
		"октября":  "Oct",
		"ноября":   "Nov",
		"декабря":  "Dec",
	}
)

func convertDateTime(dateTime string) (time.Time, error) {
	loc, err := time.LoadLocation(time_utils.TimeZoneMoscow)
	if err != nil {
		return time.Time{}, nil
	}

	timeFormatString := timeFormatStringEng
	timeOffset := time.Duration(12 * time.Hour)
	for longMonth, shortMonth := range months {
		if strings.Index(dateTime, longMonth) != -1 {
			timeFormatString = timeFormatStringRus
			timeOffset = 0
			dateTime = strings.Replace(dateTime, longMonth, shortMonth, -1)
			dateTime = strings.Replace(dateTime, "г. ", "", -1)
		}
	}

	t, err := time.ParseInLocation(timeFormatString, dateTime, loc)
	if err != nil {
		return time.Time{}, err
	}

	return t.UTC().Add(timeOffset), nil
}
