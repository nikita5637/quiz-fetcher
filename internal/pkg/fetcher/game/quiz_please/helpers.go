package quiz_please

import (
	"strconv"
	"strings"
	"time"

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
	paymentType := strings.Builder{}
	if strings.Index(game.Text, "наличные") != -1 {
		paymentType.WriteString("cash")
	}
	if strings.Index(game.Text, "карта") != -1 {
		if paymentType.Len() > 0 {
			paymentType.WriteString(",")
		}
		paymentType.WriteString("card")
	}

	// name
	name := strings.ReplaceAll(game.Name, " SPB", "")

	modelGame := model.Game{
		ExternalID:  game.ExternalID,
		LeagueID:    leagueID,
		Type:        game.GameTypeID,
		Number:      game.Number,
		Name:        name,
		PlaceID:     0,
		DateTime:    t,
		Price:       uint32(price),
		PaymentType: paymentType.String(),
		MaxPlayers:  game.MaxPlayers,
	}

	return modelGame, nil
}
