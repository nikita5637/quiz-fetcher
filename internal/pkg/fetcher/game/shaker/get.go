package shaker

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/nikita5637/quiz-fetcher/internal/pkg/logger"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	"go.uber.org/zap"
)

const (
	queryParams = "page=1&size=50&search=%7B%22city_id%22%3A%5B%22b489621b-cfb2-4aef-8c22-02daf19fa08f%22%5D%7D"
)

type location struct {
	Name        string `json:"name,omitempty"`
	Street      string `json:"street,omitempty"`
	HouseNumber string `json:"house_number,omitempty"`
}

type item struct {
	Name            string   `json:"name,omitempty"`
	Number          string   `json:"number,omitempty"`
	EventTime       string   `json:"event_time,omitempty"`
	Price           float64  `json:"price,omitempty"`
	MaxMembersCount uint32   `json:"max_members_count,omitempty"`
	Location        location `json:"location,omitempty"`
}

type data struct {
	Items []item
}

type response struct {
	Data data
}

// GetGamesList ...
func (f *Fetcher) GetGamesList(ctx context.Context) ([]model.Game, error) {
	resp, err := f.client.Get(f.url + f.gamesListPath + "?" + queryParams)
	if err != nil {
		return nil, fmt.Errorf("can't get response: %w", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body error: %w", err)
	}
	defer resp.Body.Close()

	var jsonResponse response
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error: %w", err)
	}

	games := make([]model.Game, 0, len(jsonResponse.Data.Items))
	for _, item := range jsonResponse.Data.Items {
		item.Name = strings.TrimSpace(item.Name)
		item.Location.Street = strings.TrimSpace(item.Location.Street)
		item.Location.HouseNumber = strings.TrimSpace(item.Location.HouseNumber)

		game, err := f.convertItemToGame(ctx, item)
		if err != nil {
			logger.WarnKV(ctx, "convert item to game error", zap.Error(err), zap.Reflect("item", item))
			continue
		}

		games = append(games, game)
	}

	return games, nil
}
