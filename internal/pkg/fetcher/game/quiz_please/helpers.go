package quiz_please

import (
	"strconv"
	"strings"
	"time"

	"github.com/mono83/maybe"
	"github.com/nikita5637/quiz-fetcher/internal/pkg/model"
	time_utils "github.com/nikita5637/quiz-fetcher/utils/time"
)

const (
	timeFormatString = "02.01.06 15:04"
)

func convertGameToModelGame(game game) (model.Game, error) {
	// date
	loc, err := time.LoadLocation(time_utils.TimeZoneMoscow)
	if err != nil {
		return model.Game{}, err
	}

	t, err := time.ParseInLocation(timeFormatString, game.DateTime, loc)
	if err != nil {
		return model.Game{}, err
	}

	// price
	price, err := strconv.ParseUint(game.Price[:len(game.Price)-3], 10, 32)
	if err != nil {
		return model.Game{}, err
	}

	// paymentType
	paymentTypeBuilder := strings.Builder{}
	if strings.Index(game.Text, "наличные") != -1 {
		paymentTypeBuilder.WriteString("cash")
	}
	if strings.Index(game.Text, "карта") != -1 {
		if paymentTypeBuilder.Len() > 0 {
			paymentTypeBuilder.WriteString(",")
		}
		paymentTypeBuilder.WriteString("card")
	}

	// name
	name := strings.ReplaceAll(game.Name, " SPB", "")

	paymentType := maybe.Nothing[string]()
	if v := paymentTypeBuilder.String(); v != "" {
		paymentType = maybe.Just(v)
	}

	modelGame := model.Game{
		ExternalID:  maybe.Just(game.ExternalID),
		LeagueID:    leagueID,
		Type:        game.GameTypeID,
		Number:      game.Number,
		Name:        maybe.Just(name),
		PlaceID:     0,
		DateTime:    t.UTC(),
		Price:       uint32(price),
		PaymentType: paymentType,
		MaxPlayers:  game.MaxPlayers,
		IsInMaster:  true,
	}

	return modelGame, nil
}
