package squiz

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	time_utils "github.com/nikita5637/quiz-fetcher/utils/time"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
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

var (
	gameTypesByDescription = map[string]int32{
		"Игра на общие темы. Самый популярный и массовый вариант.":                                                                                           int32(registrator.GameType_GAME_TYPE_CLASSIC),
		"Игра, предназначенная для команд без большого опыта в интеллектуальных играх.":                                                                      int32(registrator.GameType_GAME_TYPE_CLASSIC),
		"Игра на общие темы. Пакет из очень простых вопросов - идеально для новичков.":                                                                       int32(registrator.GameType_GAME_TYPE_CLASSIC),
		"Игра на общие темы. Играем одновременно один и тот же пакет вопросов во всех городах. Подведем общие итоги и наградим победителей зачетным кубком.": int32(registrator.GameType_GAME_TYPE_CLASSIC),

		"Фановая игра с раундами по мотивам популярных ТВ-шоу.":                                  int32(registrator.GameType_GAME_TYPE_THEMATIC),
		"Вопросы для взрослых, черный юмор, без цензуры. Приходите только, если уверены в себе.": int32(registrator.GameType_GAME_TYPE_THEMATIC),

		"Вопросы только на темы кино, сериалов и музыки. Зарубежное и российское, современное и классическое. Кинофаны и меломаны - вэлкам!": int32(registrator.GameType_GAME_TYPE_MOVIES_AND_MUSIC),
	}
)

func convertGameToModelGame(game game) (model.Game, error) {
	h := strings.TrimPrefix(game.Href, "#")
	externalID, err := strconv.ParseInt(h, 10, 32)
	if err != nil {
		return model.Game{}, err
	}

	gameType := gameTypesByDescription[game.Description]
	if gameType == 0 {
		return model.Game{}, errors.New("game type is 0")
	}

	if game.Number == "" {
		return model.Game{}, errors.New("empty game number")
	}

	dt := dirtyHooksForDateTime(game.DateTime)
	dateTime, err := convertDateTime(dt)
	if err != nil {
		return model.Game{}, err
	}

	price := uint64(0)
	sl := strings.Split(game.PaymentInfo, " ")
	if len(sl) > 0 {
		price, err = strconv.ParseUint(sl[0], 10, 32)
		if err != nil {
			return model.Game{}, err
		}
	}

	paymentType := ""
	if i := strings.Index(game.PaymentInfo, "наличными"); i != -1 {
		paymentType = "cash"
	}

	return model.Game{
		ExternalID:  int32(externalID),
		LeagueID:    leagueID,
		Type:        gameType,
		Number:      game.Number,
		Name:        game.Name,
		PlaceID:     0,
		DateTime:    dateTime,
		Price:       uint32(price),
		PaymentType: paymentType,
		MaxPlayers:  8,
	}, nil
}

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

func dirtyHooksForDateTime(dateTime string) string {
	switch dateTime {
	case "21января 2023 16:00":
		return "21 января 2023 16:00"
	}
	return dateTime
}

func getInfoFromHTML(text string) (string, string, error) {
	ss := strings.SplitN(text, ", ", 3)
	if len(ss) != 3 {
		return "", "", errors.New("number of elements not equal 3")
	}

	return ss[1], ss[2], nil
}

func getInfoFromPopup(ctx context.Context, html string) (string, string, string, error) {
	if html == "" {
		return "", "", "", errors.New("empty text")
	}

	ss := strings.Split(html, "<br/>")

	elements := make([]string, 0, len(ss))
	for _, s := range ss {
		if s == "" {
			continue
		}
		elements = append(elements, s)
	}

	if len(elements) != 5 {
		return "", "", "", errors.New("number of elements not equal 5")
	}

	for i := range elements {
		_ = i
		for strings.HasPrefix(elements[i], " ") {
			elements[i] = strings.TrimPrefix(elements[i], " ")
		}
		for strings.HasSuffix(elements[i], " ") {
			elements[i] = strings.TrimSuffix(elements[i], " ")
		}
	}

	dateTime := elements[0] + " " + elements[1]
	description := elements[3]
	paymentInfo := elements[4]

	return dateTime, description, paymentInfo, nil
}

func getInfoFromStrong(strong string) (string, string) {
	if strings.HasPrefix(strong, "Игра ") {
		return "", strings.TrimPrefix(strong, "Игра ")
	}
	sl := strings.Split(strong, "Игра ")
	return sl[0], sl[1]
}

func getPlaceKey(name, address string) string {
	return fmt.Sprintf("%s:%s", name, address)
}
