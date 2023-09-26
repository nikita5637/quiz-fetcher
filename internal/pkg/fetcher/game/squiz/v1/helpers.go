package squiz

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	time_utils "github.com/nikita5637/quiz-fetcher/utils/time"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
)

const (
	timeFormatString = "2 Jan 2006 15:04"
	finalGameName    = "Финал"
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

func (f *GamesFetcher) convertGameToModelGame(ctx context.Context, game game) (model.Game, error) {
	var externalID int64
	var gameType int32
	var err error
	var ret model.Game
	if game.Name != finalGameName {
		h := strings.TrimPrefix(game.Href, "#")
		externalID, err = strconv.ParseInt(h, 10, 32)
		if err != nil {
			return model.Game{}, err
		}
		ret.ExternalID = maybe.Just(int32(externalID))

		gameType, err = f.gameTypeMatchStorage.GetGameTypeByDescription(ctx, game.Description)
		if err != nil {
			return model.Game{}, fmt.Errorf("get game type by description error: %w", err)
		}

		if gameType == 0 {
			return model.Game{}, errors.New("game type is 0")
		}
		ret.Type = gameType

		price := uint64(0)
		sl := strings.Split(game.PaymentInfo, " ")
		if len(sl) > 0 {
			price, err = strconv.ParseUint(sl[0], 10, 32)
			if err != nil {
				return model.Game{}, fmt.Errorf("parse price error: %w", err)
			}
		}
		ret.Price = uint32(price)
	} else {
		ret.Type = int32(gamepb.GameType_GAME_TYPE_CLOSED)
	}

	if game.Number == "" {
		return model.Game{}, errors.New("empty game number")
	}

	ret.Number = game.Number
	if i := strings.Index(game.Number, "("); i != -1 {
		ret.Number = game.Number[:i]
	}

	dt := dirtyHooksForDateTime(game.DateTime)
	dateTime, err := convertDateTime(dt)
	if err != nil {
		return model.Game{}, fmt.Errorf("convert date time error: %w", err)
	}
	ret.DateTime = dateTime

	paymentType := ""
	if i := strings.Index(game.PaymentInfo, "наличными"); i != -1 {
		paymentType = "cash"
	}

	ret.PaymentType = maybe.Nothing[string]()
	if paymentType != "" {
		ret.PaymentType = maybe.Just(paymentType)
	}

	ret.Name = maybe.Nothing[string]()
	if game.Name != "" {
		ret.Name = maybe.Just(game.Name)
	}

	ret.LeagueID = leagueID
	ret.PlaceID = 0
	ret.MaxPlayers = maxPlayers
	ret.IsInMaster = true

	return ret, nil
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

func getInfoFromCommonGamePopup(ctx context.Context, html string) (string, string, string, error) {
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

func getInfoFromFinalGamePopup(ctx context.Context, html string) (string, string, error) {
	if html == "" {
		return "", "", errors.New("empty text")
	}

	ss := strings.Split(html, "<br/>")

	elements := make([]string, 0, len(ss))
	for _, s := range ss {
		if s == "" {
			continue
		}
		elements = append(elements, s)
	}

	if len(elements) != 6 {
		return "", "", errors.New("number of elements not equal 5")
	}

	for i := range elements {
		_ = i
		for strings.HasPrefix(elements[i], " ") {
			elements[i] = strings.TrimPrefix(elements[i], " ")
		}
		for strings.HasPrefix(elements[i], "- ") {
			elements[i] = strings.TrimPrefix(elements[i], "- ")
		}
		for strings.HasSuffix(elements[i], " ") {
			elements[i] = strings.TrimSuffix(elements[i], " ")
		}
	}

	dateTime := elements[5]
	year := strconv.Itoa(time_utils.TimeNow().Year())
	for _, day := range []string{"(пн),", "(вт),", "(ср),", "(чт),", "(пт),", "(сб),", "(вс),"} {
		dateTime = strings.Replace(dateTime, day, year, 1)
	}

	return dateTime, "наличными", nil
}

func getInfoFromStrong(strong string) (string, string) {
	if strings.HasPrefix(strong, "Игра ") {
		return "", strings.TrimPrefix(strong, "Игра ")
	}

	if strings.HasPrefix(strong, "Финал. Сезон ") {
		sl := strings.Split(strong, ". ")
		if len(sl) == 2 {
			return sl[0], sl[1]
		}

		return "", ""
	}

	sl := strings.Split(strong, "Игра ")
	if len(sl) == 2 {
		return sl[0], sl[1]
	}

	return "", ""
}
