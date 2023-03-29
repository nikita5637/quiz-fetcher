package sixty_seconds

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/logger"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"
)

// GetGamesList ...
func (f *GamesFetcher) GetGamesList(ctx context.Context) ([]model.Game, error) {
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
				logger.WarnKV(ctx, "can't parse externalID", "error", err, "gameInfoPath", gameInfoPath)
				return
			}
			g.ExternalID = externalID

			h5Text := h5.Text()
			name := getName(h5Text)
			if name == "" {
				logger.WarnKV(ctx, "can't parse game name", "text", h5.Text())
				return
			}
			g.Name = name

			number := getNumber(h5Text)
			if number == "" {
				logger.WarnKV(ctx, "can't parse game number", "text", h5.Text())
				return
			}
			g.Number = number

			if strings.HasPrefix(number, "#") {
				g.Type = int32(registrator.GameType_GAME_TYPE_CLASSIC)
			} else if number == "Финал" {
				g.Type = int32(registrator.GameType_GAME_TYPE_CLASSIC)
			}

			if g.Type == 0 {
				logger.WarnKV(ctx, "can't parse game type", "number", number)
				return
			}

			dateTime, err := f.getDateTime(ctx, gameInfoPath)
			if err != nil {
				logger.Warnf(ctx, "can't parse game dateTime: %s", err.Error())
				return
			}
			g.DateTime = dateTime

			g.PaymentType = "cash"
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
				logger.WarnKV(ctx, "can't parse place", "error", err, "place", placeStr)
				return
			}
			g.PlaceID = placeID

			price, err := getPrice(priceStr)
			if err != nil {
				logger.WarnKV(ctx, "can't parse price", "error", err, "price", priceStr)
				return
			}
			g.Price = price

			games = append(games, g)
		})
	}

	return games, nil
}

func (f *GamesFetcher) getDateTime(ctx context.Context, gameInfoPath string) (time.Time, error) {
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

func (f *GamesFetcher) getPlaceID(ctx context.Context, place string) (int32, error) {
	sl := strings.Split(place, " - ")
	if len(sl) != 2 {
		return 0, errors.New("can't parse place string")
	}

	name := strings.TrimPrefix(sl[0], "\u00a0")
	address := sl[1]

	placeResp, err := f.registratorServiceClient.GetPlaceByNameAndAddress(ctx, &registrator.GetPlaceByNameAndAddressRequest{
		Address: address,
		Name:    name,
	})
	if err != nil {
		return 0, err
	}

	return placeResp.GetPlace().GetId(), nil
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
	s := strings.TrimSuffix(strings.TrimPrefix(price, "\u00a0"), " руб. с команды")
	p, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return uint32(p), nil
}
