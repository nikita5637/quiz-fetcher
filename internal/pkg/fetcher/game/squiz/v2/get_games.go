package squiz

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/logger"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
)

// Characteristic ...
type Characteristic struct {
	Title string
	Value string
}

// Product ...
type Product struct {
	Title           string
	Characteristics []Characteristic
	Descr           string
	Price           string
	Text            string
}

// StoreResponse ...
type StoreResponse struct {
	Products []Product
}

// GetGamesList ...
func (f *GamesFetcher) GetGamesList(ctx context.Context) ([]model.Game, error) {
	req, err := http.NewRequest(http.MethodGet, f.url+f.gamesListPath+gamesListQuery, nil)
	if err != nil {
		return nil, fmt.Errorf("create new http request error: %w", err)
	}

	req.Header.Add("Host", "store.tildacdn.com")
	req.AddCookie(&http.Cookie{
		Name:  "__ddg1_",
		Value: "0l5CSRIalyK3p1M8JcMp",
	})
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

	resp, err := f.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can't get games list: %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body error: %w", err)
	}
	defer resp.Body.Close()

	storeResponse := &StoreResponse{}
	if err := json.Unmarshal(body, &storeResponse); err != nil {
		return nil, fmt.Errorf("json unmarshal error: %w", err)
	}

	modelGames := make([]model.Game, 0)
	for _, product := range storeResponse.Products {
		gameFormat := getGameFormat(product.Characteristics)
		if gameFormat == "Онлайн" {
			logger.InfoKV(ctx, "skip online game")
			continue
		}

		game, err := f.getGameFromProduct(ctx, product)
		if err != nil {
			logger.WarnKV(ctx, "get game from product error", "error", err.Error(), "product", product)
			continue
		}

		modelGames = append(modelGames, game)
	}

	return modelGames, nil
}

func (f *GamesFetcher) getGameFromProduct(ctx context.Context, product Product) (model.Game, error) {
	gameDescription := getGameDescription(product.Text)
	if gameDescription == "" {
		return model.Game{}, errors.New("empty description")
	}

	gameName := getGameName(product.Text)

	gameType, err := f.getGameType(ctx, gameName, gameDescription)
	if err != nil {
		return model.Game{}, fmt.Errorf("get game type error: %w", err)
	}

	gameDate := getGameDate(product.Text)
	if gameDate == "" {
		return model.Game{}, errors.New("empty game date")
	}

	gameTime := getGameTime(product.Title)
	if gameTime == "" {
		return model.Game{}, errors.New("empty game time")
	}

	gameDateTime, err := convertDateTime(gameDate + " " + gameTime)
	if err != nil {
		return model.Game{}, fmt.Errorf("convert game datetime error: %w", err)
	}

	gamePlaceAddress := getGamePlaceAddress(product.Text)
	if gamePlaceAddress == "" {
		return model.Game{}, errors.New("empty game place address")
	}

	gamePlaceName := getGamePlaceName(product.Text)
	if gamePlaceName == "" {
		return model.Game{}, errors.New("empty game place name")
	}

	place, err := f.placeStorage.GetPlaceByNameAndAddress(ctx, gamePlaceName, gamePlaceAddress)
	if err != nil {
		return model.Game{}, fmt.Errorf("get place by name and address error: %w", err)
	}

	price, err := strconv.ParseFloat(product.Price, 64)
	if err != nil {
		return model.Game{}, fmt.Errorf("parse price error: %w", err)
	}

	name := maybe.Nothing[string]()
	if gameName != "" {
		name = maybe.Just(gameName)
	}

	return model.Game{
		ExternalID:  maybe.Nothing[int32](),
		LeagueID:    leagueID,
		Type:        gameType,
		Number:      product.Descr,
		Name:        name,
		PlaceID:     int32(place.ExternalID),
		DateTime:    gameDateTime,
		Price:       uint32(price),
		PaymentType: maybe.Just("cash"),
		MaxPlayers:  maxPlayers,
		IsInMaster:  true,
	}, nil
}

func (f *GamesFetcher) getGameType(ctx context.Context, gameName, description string) (int32, error) {
	if gameName == finalGameName {
		return int32(gamepb.GameType_GAME_TYPE_CLOSED), nil
	}

	gameType, err := f.gameTypeMatchStorage.GetGameTypeByDescription(ctx, description)
	if err != nil {
		return 0, err
	}

	return gameType, nil
}
