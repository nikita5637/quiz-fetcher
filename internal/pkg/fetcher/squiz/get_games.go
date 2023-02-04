package squiz

import (
	"context"
	"fmt"
	"strings"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/logger"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	"github.com/nikita5637/quiz-registrator-api/pkg/pb/registrator"

	"github.com/PuerkitoBio/goquery"
)

type game struct {
	Href string // external ID

	Description string
	Name        string
	Number      string

	PlaceName    string // placeID
	PlaceAddress string // placeID

	DateTime    string
	PaymentInfo string
	MaxPlayers  uint32
}

// GetGamesList ...
func (f *GamesFetcher) GetGamesList(ctx context.Context) ([]model.Game, error) {
	resp, err := f.client.Get(f.url + f.gamesListPath)
	if err != nil {
		return nil, fmt.Errorf("can't get games list: %w", err)
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("can't read document: %w", err)
	}

	games := make(map[string]*game, 12)

	doc.Find(".t-item").Each(func(_ int, item *goquery.Selection) {
		parentDiv := item.Children().First()
		parentA := parentDiv.Find("a").First()
		parentTText := parentDiv.Find(".t-text").First()

		strongText := parentTText.Find("strong").First().Text()
		gameName, gameNumber := getInfoFromStrong(strongText)

		parentTText.Children().Empty()

		href, exists := parentA.Attr("href")
		if exists {
			if strings.HasPrefix(href, "#") {
				text := parentTText.Text()
				place, address, err := getInfoFromHTML(text)
				if err != nil {
					logger.WarnKV(ctx, "can't parse a game's info", "text", text, "error", err)
					return
				}

				games[href] = &game{
					Href:         href,
					Name:         gameName,
					Number:       gameNumber,
					PlaceName:    place,
					PlaceAddress: address,
				}
			}
		}
	})

	doc.Find("div").Each(func(_ int, div *goquery.Selection) {
		value, exists := div.Attr("data-tooltip-hook")
		if !exists {
			return
		}

		game, ok := games[value]
		if !ok {
			return
		}

		description := div.Find(".t-descr")
		description.Find("strong").Remove()
		description.Find("a").Remove()
		description.Find("span").Remove()

		html, _ := description.Html()
		dateTime, gameDescription, paymentInfo, err := getInfoFromPopup(ctx, html)
		if err != nil {
			logger.WarnKV(ctx, "can't parse a game's popup", "html", html, "error", err)
			return
		}

		game.Description = gameDescription
		game.DateTime = dateTime
		game.PaymentInfo = paymentInfo
	})

	modelGames := make([]model.Game, 0, len(games))
	for _, game := range games {
		modelGame, err := f.convertGameToModelGame(ctx, *game)
		if err != nil {
			logger.WarnKV(ctx, "can't convert game", "error", err.Error(), "game", game)
			continue
		}

		placeID, ok := f.placesCache[getPlaceKey(game.PlaceName, game.PlaceAddress)]
		if ok {
			logger.DebugKV(ctx, "place found in cache", "name", game.PlaceName, "address", game.PlaceAddress, "id", placeID)
		} else {
			logger.DebugKV(ctx, "place not found in cache", "name", game.PlaceName, "address", game.PlaceAddress)
			placeResp, err := f.registratorServiceClient.GetPlaceByNameAndAddress(ctx, &registrator.GetPlaceByNameAndAddressRequest{
				Name:    game.PlaceName,
				Address: game.PlaceAddress,
			})
			if err != nil {
				logger.Warnf(ctx, "can't get place ID: %s", err.Error())
				continue
			}

			placeID = placeResp.GetPlace().GetId()
			f.placesCache[getPlaceKey(game.PlaceName, game.PlaceAddress)] = placeID
			logger.DebugKV(ctx, "place stored in cache", "name", game.PlaceName, "address", game.PlaceAddress, "id", placeID)
		}
		modelGame.PlaceID = placeID

		modelGames = append(modelGames, modelGame)
	}

	return modelGames, nil
}
