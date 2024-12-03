package sixty_seconds

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/logger"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	"go.uber.org/zap"
)

const openLeague = "Открытая лига"
const firstLeague = "Первая лига"
const openLeagueFinal = "Открытая лига Финал"
const final = "Финал"

// GetGamesList ...
func (f *Fetcher) GetGamesList(ctx context.Context) ([]model.Game, error) {
	resp, err := f.client.Get(f.url + f.schedulePath)
	if err != nil {
		return nil, fmt.Errorf("failed to get response: %w", err)
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to create document from response: %w", err)
	}

	games := make([]model.Game, 0)
	doc.Find(".container").Each(func(i int, c *goquery.Selection) {
		if !c.HasClass("py-3") {
			return
		}

		if id, ok := c.Parent().Attr("id"); ok {
			if id == "header" {
				return
			}
		}

		game := model.Game{
			LeagueID: leagueID,
		}

		h5 := c.Find("h5").First()
		a := h5.Find("a").First()

		var gameInfoPath string
		if value, exists := a.Attr("href"); exists {
			gameInfoPath = value
		}

		externalID, err := getExternalID(gameInfoPath)
		if err != nil {
			logger.WarnKV(ctx, "failed to parse externalID", zap.Error(err), zap.String("game_info_path", gameInfoPath))
			return
		}
		game.ExternalID = maybe.Just(externalID)

		h5Text := h5.Text()
		name := getName(h5Text)
		if name == "" {
			logger.WarnKV(ctx, "game name is empty", zap.String("text", h5.Text()))
			return
		}

		if name == openLeague {
			if !f.needToFetchOpenLeague {
				return
			}
		} else if name == firstLeague {
			if !f.needToFetchFirstLeague {
				return
			}
		} else {
			return
		}

		game.Name = maybe.Just(name)

		number := getNumber(h5Text)
		if number == "" {
			logger.WarnKV(ctx, "game number is empty", zap.String("text", h5.Text()))
			return
		}
		game.Number = number

		if strings.HasPrefix(number, "#") {
			game.Type = int32(gamepb.GameType_GAME_TYPE_CLASSIC)
		} else if number == final || number == "Финал сезона" {
			game.Type = int32(gamepb.GameType_GAME_TYPE_CLASSIC)
		}

		if game.Type == 0 {
			logger.WarnKV(ctx, "game type is empty", zap.String("number", number))
			return
		}

		dateTime, err := f.getDateTime(ctx, gameInfoPath)
		if err != nil {
			logger.WarnKV(ctx, "failed to get game date and time", zap.Error(err))
			return
		}
		game.DateTime = dateTime

		game.PaymentType = maybe.Just("cash")
		game.MaxPlayers = 6

		var placeStr, priceStr string
		c.Find("tr").Each(func(i int, tr *goquery.Selection) {
			switch i {
			case 2:
				placeStr = tr.Text()
			case 3:
				priceStr = tr.Text()
			}
		})

		placeID, err := f.getPlaceID(ctx, strings.TrimSpace(placeStr))
		if err != nil {
			logger.WarnKV(ctx, "failed to get place ID", zap.Error(err), zap.String("place", placeStr))
			return
		}
		game.PlaceID = int32(placeID)

		price, err := getPrice(strings.TrimSpace(priceStr))
		if err != nil {
			logger.WarnKV(ctx, "parsing price error", zap.Error(err), zap.String("price", priceStr))
			return
		}
		game.Price = price
		game.IsInMaster = true

		games = append(games, game)
	})

	return games, nil
}

func (f *Fetcher) getDateTime(_ context.Context, gameInfoPath string) (time.Time, error) {
	req, err := http.NewRequest("GET", f.url+gameInfoPath, nil)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to create new request: %w", err)
	}

	req.Header.Add("Accept-Language", "en-GB,en;q=0.9")

	resp, err := f.client.Do(req)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to do a request: %w", err)
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return time.Time{}, fmt.Errorf("can't read document: %w", err)
	}

	dateTime := ""
	doc.Find(".l-blue").Each(func(_ int, s *goquery.Selection) {
		sl := strings.Split(s.Text(), " /\n")
		dateTime = sl[len(sl)-1]
	})

	return convertDateTime(strings.TrimSpace(dateTime))
}

func (f *Fetcher) getPlaceID(ctx context.Context, place string) (int, error) {
	name := ""
	address := ""
	sl := strings.Split(place, " - ")
	if len(sl) != 2 {
		sl = strings.Split(place, ", ")
		if len(sl) < 2 {
			return 0, errors.New("can't parse place string")
		}

		name = strings.TrimSpace(sl[0])
		address = strings.Join(sl[1:], ", ")
	} else {
		name = strings.TrimSpace(sl[0])
		address = sl[1]
	}

	dbPlace, err := f.placeStorage.GetPlaceByNameAndAddress(ctx, name, address)
	if err != nil {
		return 0, err
	}

	return dbPlace.ExternalID, nil
}

func getExternalID(path string) (int32, error) {
	s := strings.TrimSuffix(strings.TrimPrefix(path, "/quizgames/game/"), "/")
	id, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return int32(id), nil
}

func getName(text string) string {
	s := strings.Split(text, " | ")
	if len(s) == 2 {
		return strings.TrimSpace(s[0])
	}

	if text == openLeagueFinal {
		return openLeague
	}

	return ""
}

func getNumber(text string) string {
	s := strings.Split(text, " | ")
	if len(s) == 2 {
		ret := s[1]
		if strings.HasPrefix(s[1], "Игра ") {
			ret = strings.TrimPrefix(s[1], "Игра ")
		}

		return strings.TrimSpace(ret)
	}

	if text == openLeagueFinal {
		return final
	}

	return ""
}

func getPrice(price string) (uint32, error) {
	if strings.HasSuffix(price, " руб. с команды + депозит") {
		price = strings.TrimSuffix(strings.TrimPrefix(price, "\u00a0"), " руб. с команды + депозит")
	} else if strings.HasSuffix(price, " руб. с команды") {
		price = strings.TrimSuffix(strings.TrimPrefix(price, "\u00a0"), " руб. с команды")
	} else if strings.HasSuffix(price, " руб. с человека") {
		price = strings.TrimSuffix(strings.TrimPrefix(price, "\u00a0"), " руб. с человека")
	}

	p, err := strconv.Atoi(price)
	if err != nil {
		return 0, err
	}

	return uint32(p), nil
}
