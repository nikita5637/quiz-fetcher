package sixty_seconds

import (
	"errors"
	"strings"
	"time"

	time_utils "github.com/nikita5637/quiz-fetcher/utils/time"
)

const (
	timeFormatString = "2 Jan 2006 15:04"
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

	ret := ""
	for longMonth, shortMonth := range months {
		if index := strings.Index(dateTime, longMonth); index != -1 {
			ret = strings.Replace(dateTime, longMonth, shortMonth, -1)
		}
	}

	if ret == "" {
		return time.Time{}, errors.New("can't convert time")
	}

	t, err := time.ParseInLocation(timeFormatString, ret, loc)
	if err != nil {
		return time.Time{}, err
	}

	return t.UTC(), nil
}
