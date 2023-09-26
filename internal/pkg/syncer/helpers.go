package syncer

import (
	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	gamepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/game"
	leaguepb "github.com/nikita5637/quiz-registrator-api/pkg/pb/league"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func convertModelGameToProtoGame(game model.Game) *gamepb.Game {
	ret := &gamepb.Game{
		Id:         game.ID,
		LeagueId:   leaguepb.LeagueID(game.LeagueID),
		Type:       gamepb.GameType(game.Type),
		Number:     game.Number,
		PlaceId:    game.PlaceID,
		Date:       timestamppb.New(game.DateTime),
		Price:      game.Price,
		MaxPlayers: game.MaxPlayers,
		IsInMaster: game.IsInMaster,
	}

	if externalID, isPresent := game.ExternalID.Get(); isPresent {
		ret.ExternalId = wrapperspb.Int32(externalID)
	}

	if name, isPresent := game.Name.Get(); isPresent {
		ret.Name = wrapperspb.String(name)
	}

	if paymentType, isPresent := game.PaymentType.Get(); isPresent {
		ret.PaymentType = wrapperspb.String(paymentType)
	}

	return ret
}

func convertProtoGameToModelGame(game *gamepb.Game) model.Game {
	ret := model.Game{
		ID:          game.GetId(),
		ExternalID:  maybe.Nothing[int32](),
		LeagueID:    int32(game.LeagueId),
		Type:        int32(game.GetType()),
		Number:      game.GetNumber(),
		Name:        maybe.Nothing[string](),
		PlaceID:     game.GetPlaceId(),
		DateTime:    game.GetDate().AsTime(),
		Price:       game.GetPrice(),
		PaymentType: maybe.Nothing[string](),
		MaxPlayers:  game.GetMaxPlayers(),
		IsInMaster:  game.GetIsInMaster(),
	}

	if externalID := game.GetExternalId(); externalID != nil {
		ret.ExternalID = maybe.Just(externalID.GetValue())
	}

	if name := game.GetName(); name != nil {
		ret.Name = maybe.Just(name.GetValue())
	}

	if paymentType := game.GetPaymentType(); paymentType != nil {
		ret.PaymentType = maybe.Just(paymentType.GetValue())
	}

	return ret
}
