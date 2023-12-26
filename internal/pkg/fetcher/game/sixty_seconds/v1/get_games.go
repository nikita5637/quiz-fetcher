package sixty_seconds

import (
	"context"
	"errors"
	"fmt"
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

// GetGamesList ...
func (f *Fetcher) GetGamesList(ctx context.Context) ([]model.Game, error) {
	games := []model.Game{}
	for _, gamesListPath := range []string{f.openLeagueGamesListPath, f.firstLeagueGamesListPath} {
		if gamesListPath == "" {
			continue
		}

		resp, err := f.client.Get(f.url + gamesListPath)
		if err != nil {
			return nil, fmt.Errorf("can't get response: %w", err)
		}

		doc, err := goquery.NewDocumentFromResponse(resp)
		if err != nil {
			return nil, fmt.Errorf("can't read document: %w", err)
		}

		doc.Find(".container").Each(func(i int, c *goquery.Selection) {
			if !c.HasClass("py-3") {
				return
			}

			g := model.Game{
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
				logger.WarnKV(ctx, "parsing externalID error", zap.Error(err), zap.String("game_info_path", gameInfoPath))
				return
			}
			g.ExternalID = maybe.Just(externalID)

			h5Text := h5.Text()
			name := getName(h5Text)
			if name == "" {
				logger.WarnKV(ctx, "game name is empty", zap.String("text", h5.Text()))
				return
			}
			g.Name = maybe.Just(name)

			number := getNumber(h5Text)
			if number == "" {
				logger.WarnKV(ctx, "game number is empty", zap.String("text", h5.Text()))
				return
			}
			g.Number = number

			if strings.HasPrefix(number, "#") {
				g.Type = int32(gamepb.GameType_GAME_TYPE_CLASSIC)
			} else if number == "Финал" || number == "Финал сезона" {
				g.Type = int32(gamepb.GameType_GAME_TYPE_CLASSIC)
			}

			if g.Type == 0 {
				logger.WarnKV(ctx, "game type is empty", zap.String("number", number))
				return
			}

			dateTime, err := f.getDateTime(ctx, gameInfoPath)
			if err != nil {
				logger.WarnKV(ctx, "parsing game date and time error", zap.Error(err))
				return
			}
			g.DateTime = dateTime

			g.PaymentType = maybe.Just("cash")
			g.MaxPlayers = 6

			var placeStr, priceStr string
			c.Find("tr").Each(func(i int, tr *goquery.Selection) {
				switch i {
				case 2:
					placeStr = tr.Text()
				case 3:
					priceStr = tr.Text()
				}
			})

			placeID, err := f.getPlaceID(ctx, placeStr)
			if err != nil {
				logger.WarnKV(ctx, "parsing place error", zap.Error(err), zap.String("place", placeStr))
				return
			}
			g.PlaceID = int32(placeID)

			price, err := getPrice(priceStr)
			if err != nil {
				logger.WarnKV(ctx, "parsing price error", zap.Error(err), zap.String("price", priceStr))
				return
			}
			g.Price = price
			g.IsInMaster = true

			games = append(games, g)
		})
	}

	return games, nil
}

func (f *Fetcher) getDateTime(ctx context.Context, gameInfoPath string) (time.Time, error) {
	resp, err := f.client.Get(f.url + gameInfoPath)
	if err != nil {
		return time.Time{}, fmt.Errorf("can't get response: %w", err)
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return time.Time{}, fmt.Errorf("can't read document: %w", err)
	}

	dateTime := ""
	doc.Find(".l-blue").Each(func(_ int, s *goquery.Selection) {
		sl := strings.Split(s.Text(), " / ")
		dateTime = sl[len(sl)-1]
	})

	return convertDateTime(dateTime)
}

func (f *Fetcher) getPlaceID(ctx context.Context, place string) (int, error) {
	sl := strings.Split(place, " - ")
	if len(sl) != 2 {
		return 0, errors.New("can't parse place string")
	}

	name := strings.TrimSpace(sl[0])
	address := sl[1]

	dbPlace, err := f.placeStorage.GetPlaceByNameAndAddress(ctx, name, address)
	if err != nil {
		return 0, err
	}

	return dbPlace.ExternalID, nil
}

func getName(text string) string {
	s := strings.Split(text, " | ")
	if len(s) == 2 {
		return s[0]
	}

	return ""
}

func getNumber(text string) string {
	s := strings.Split(text, " | ")
	if len(s) == 2 {
		if strings.HasPrefix(s[1], "Игра ") {
			return strings.TrimPrefix(s[1], "Игра ")
		}

		return s[1]
	}

	return ""
}

func getExternalID(path string) (int32, error) {
	s := strings.TrimSuffix(strings.TrimPrefix(path, "/game/"), "/")
	id, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return int32(id), nil
}

func getPrice(price string) (uint32, error) {
	s := price
	if strings.HasSuffix(price, " руб. с команды") {
		s = strings.TrimSuffix(strings.TrimPrefix(price, "\u00a0"), " руб. с команды")
	} else if strings.HasSuffix(price, " руб. с человека") {
		s = strings.TrimSuffix(strings.TrimPrefix(price, "\u00a0"), " руб. с человека")
	}

	p, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return uint32(p), nil
}
