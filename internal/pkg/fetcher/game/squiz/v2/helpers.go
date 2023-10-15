package squiz

import (
	"bytes"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	time_utils "github.com/nikita5637/quiz-fetcher/utils/time"
)

const (
	timeFormatString = "2006-01-02 15:04"
	finalGameName    = "Финал"
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

	return t.UTC(), nil
}

func getGameDate(text string) string {
	sl := strings.Split(text, "<br />")
	if len(sl) == 0 {
		return ""
	}

	ret := ""
	for _, v := range sl {
		if strings.HasPrefix(v, "Дата: ") {
			ret = strings.TrimPrefix(v, "Дата: ")
		}
	}

	ret = strings.TrimPrefix(ret, `"`)
	ret = strings.TrimSuffix(ret, `"`)

	return ret
}

func getGameDescription(text string) string {
	sl := strings.Split(text, "<br />")
	if len(sl) == 0 {
		return ""
	}

	ret := ""
	for _, v := range sl {
		if strings.HasPrefix(v, "Текст попапа: ") {
			ret = strings.TrimPrefix(v, "Текст попапа: ")
		}
	}

	ret = strings.TrimPrefix(ret, `"`)
	ret = strings.TrimSuffix(ret, `"`)

	return ret
}

func getGameFormat(characteristics []Characteristic) string {
	for _, characharacteristic := range characteristics {
		if characharacteristic.Title == "Формат игры" {
			return characharacteristic.Value
		}
	}

	return ""
}

func getGameName(text string) string {
	sl := strings.Split(text, "<br />")
	if len(sl) == 0 {
		return ""
	}

	ret := strings.TrimPrefix(sl[0], "Описание: ")
	ret = strings.TrimPrefix(ret, `"`)
	ret = strings.TrimSuffix(ret, `"`)

	return ret
}

func getGamePlaceAddress(text string) string {
	sl := strings.Split(text, "<br />")
	if len(sl) == 0 {
		return ""
	}

	for _, v := range sl {
		if strings.HasPrefix(v, "Локация: ") {
			s := strings.TrimPrefix(v, "Локация: ")
			r := bytes.NewBuffer([]byte(s))
			doc, err := goquery.NewDocumentFromReader(r)
			if err != nil {
				return ""
			}
			txt := doc.Text()

			sl2 := strings.Split(txt, "\", \"")
			if len(sl2) < 2 {
				return ""
			}

			return sl2[1]
		}
	}

	return ""
}

func getGamePlaceName(text string) string {
	sl := strings.Split(text, "<br />")
	if len(sl) == 0 {
		return ""
	}

	for _, v := range sl {
		if strings.HasPrefix(v, "Локация: ") {
			s := strings.TrimPrefix(v, "Локация: ")
			r := bytes.NewBuffer([]byte(s))
			doc, err := goquery.NewDocumentFromReader(r)
			if err != nil {
				return ""
			}
			txt := doc.Text()

			sl2 := strings.Split(txt, "\", \"")
			if len(sl2) == 0 {
				return ""
			}

			ret := strings.TrimPrefix(sl2[0], `"`)
			ret = strings.TrimPrefix(ret, " ")
			ret = strings.TrimSuffix(ret, " ")

			return ret
		}
	}

	return ""
}

func getGameTime(title string) string {
	sl := strings.Split(title, " :: ")
	if len(sl) < 2 {
		return ""
	}

	return sl[1]
}
