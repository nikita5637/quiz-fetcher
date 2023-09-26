package time

import "time"

const (
	// TimeZoneMoscow ...
	TimeZoneMoscow = "Europe/Moscow"
)

var (
	// TimeNow returns time in UTC
	TimeNow = time.Now().UTC
)

// ConvertTime converts string to time in UTC
func ConvertTime(str string) time.Time {
	timeFormat := "2006-01-02 15:04"

	t, err := time.Parse(timeFormat, str)
	if err != nil {
		return time.Time{}
	}

	return t
}
