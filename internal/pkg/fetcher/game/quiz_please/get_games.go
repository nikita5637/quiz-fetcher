package quiz_please

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/logger"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	"go.uber.org/zap"
)

type game struct {
	ExternalID int32  `json:"gameId"`
	GameTypeID int32  `json:"game_type_id"`
	GameType   int32  `json:"game_type"`
	Name       string `json:"nameGame"`
	Number     string `json:"numberGame"`

	PlaceName    string `json:"place"`   // placeID
	PlaceAddress string `json:"address"` // placeID

	DateTime string `json:"datetime"`
	Price    string `json:"price"`

	Text       string `json:"text"` // paymentType
	MaxPlayers uint32 `json:"max_players"`
}

// GetGamesList ...
func (f *Fetcher) GetGamesList(ctx context.Context) ([]model.Game, error) {
	gameIDs, err := f.getGameIDs(ctx)
	if err != nil {
		return nil, err
	}

	return f.getGames(ctx, gameIDs), nil
}

func (f *Fetcher) getGame(gameID int64) (game, error) {
	url := fmt.Sprintf(f.url+f.gameInfoPathFormat, gameID)
	resp, err := f.client.Get(url)
	if err != nil {
		return game{}, fmt.Errorf("can't get game info: %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return game{}, fmt.Errorf("can't read body: %w", err)
	}
	defer resp.Body.Close()

	g := game{}
	err = json.Unmarshal(body, &g)
	if err != nil {
		return game{}, fmt.Errorf("can't unmarshal body: %w", err)
	}

	return g, nil
}

// skips games if
// error while getting game occured
// error while convert game to model.Game occured
// error while get place id occured
func (f *Fetcher) getGames(ctx context.Context, gameIDs []int64) []model.Game {
	games := make([]model.Game, 0, len(gameIDs))
	countOnlineGames := 0
	for _, gameID := range gameIDs {
		g, err := f.getGame(gameID)
		if err != nil {
			logger.WarnKV(ctx, "getting game error", zap.Error(err))
			continue
		}

		modelGame, err := convertGameToModelGame(g)
		if err != nil {
			logger.WarnKV(ctx, "converting game error", zap.Error(err))
			continue
		}

		if g.GameType == 1 {
			countOnlineGames++
			continue
		}

		place, err := f.placeStorage.GetPlaceByNameAndAddress(ctx, g.PlaceName, g.PlaceAddress)
		if err != nil {
			logger.WarnKV(ctx, "getting place by name and address error", zap.Error(err), zap.String("place_name", g.PlaceName), zap.String("place_address", g.PlaceAddress))
			continue
		}

		modelGame.PlaceID = int32(place.ExternalID)
		games = append(games, modelGame)
	}

	if countOnlineGames > 0 {
		logger.Infof(ctx, "skipped %d online games", countOnlineGames)
	}

	return games
}

func (f *Fetcher) getGameIDs(ctx context.Context) ([]int64, error) {
	resp, err := f.client.Get(f.url + f.gamesListPath)
	if err != nil {
		return nil, fmt.Errorf("can't get response: %w", err)
	}

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("can't read document: %w", err)
	}

	gameIDs := make([]int64, 0)
	doc.Find("#w1").Each(func(_ int, w1 *goquery.Selection) {
		w1.Find(".schedule-column").Each(func(_ int, game *goquery.Selection) {
			if value, exists := game.Attr("id"); exists {
				if gameID, err := strconv.ParseInt(value, 10, 64); err != nil {
					logger.WarnKV(ctx, "parsing gameID error", zap.Error(err))
				} else {
					logger.Debugf(ctx, "game ID %d found", gameID)
					gameIDs = append(gameIDs, gameID)
				}
			}
		})
	})

	return gameIDs, nil
}
