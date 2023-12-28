package shaker

import (
	"context"
	"fmt"
	"time"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	time_utils "github.com/nikita5637/quiz-fetcher/utils/time"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
)

const (
	defaultGameType  = int32(gamepb.GameType_GAME_TYPE_THEMATIC)
	timeFormatString = "2006-01-02T15:04:00"
)

func (f *Fetcher) convertItemToGame(ctx context.Context, item item) (model.Game, error) {
	ret := model.Game{
		ExternalID:  maybe.Nothing[int32](),
		LeagueID:    leagueID,
		Number:      item.Number,
		Name:        maybe.Nothing[string](),
		Price:       uint32(item.Price),
		PaymentType: maybe.Just("cash"),
		MaxPlayers:  item.MaxMembersCount,
		IsInMaster:  true,
	}

	gameType, err := f.gameTypeMatchStorage.GetGameTypeByName(ctx, item.Name)
	if err != nil {
		return model.Game{}, fmt.Errorf("getting gameType by name error: %w", err)
	}

	if gameType == 0 {
		gameType = defaultGameType
	}

	ret.Type = gameType

	if item.Name != "" {
		ret.Name = maybe.Just(item.Name)
	}

	place, err := f.placeStorage.GetPlaceByNameAndAddress(ctx, item.Location.Name, fmt.Sprintf("%s %s", item.Location.Street, item.Location.HouseNumber))
	if err != nil {
		return model.Game{}, fmt.Errorf("getting place by name and address error: %w", err)
	}
	ret.PlaceID = int32(place.ExternalID)

	dateTime, err := convertDateTime(item.EventTime)
	if err != nil {
		return model.Game{}, fmt.Errorf("converting date error: %w", err)
	}
	ret.DateTime = dateTime

	return ret, nil
}

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
