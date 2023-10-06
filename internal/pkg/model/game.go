package model

import (
	"encoding/json"
	"time"

	"github.com/mono83/maybe"
	maybejson "github.com/mono83/maybe/json"
)

// Game ...
type Game struct {
	ID          int32
	ExternalID  maybe.Maybe[int32]
	LeagueID    int32
	Type        int32
	Number      string
	Name        maybe.Maybe[string]
	PlaceID     int32
	DateTime    time.Time
	Price       uint32
	PaymentType maybe.Maybe[string]
	MaxPlayers  uint32
	IsInMaster  bool
}

// MarshalJSON ...
func (g Game) MarshalJSON() ([]byte, error) {
	type wrapperGame struct {
		ID          int32
		ExternalID  maybejson.Maybe[int32]
		LeagueID    int32
		Type        int32
		Number      string
		Name        maybejson.Maybe[string]
		PlaceID     int32
		DateTime    time.Time
		Price       uint32
		PaymentType maybejson.Maybe[string]
		MaxPlayers  uint32
	}

	wg := wrapperGame{
		ID:          g.ID,
		ExternalID:  maybejson.Wrap(g.ExternalID),
		LeagueID:    g.LeagueID,
		Type:        g.Type,
		Number:      g.Number,
		Name:        maybejson.Wrap(g.Name),
		PlaceID:     g.PlaceID,
		DateTime:    g.DateTime,
		Price:       g.Price,
		PaymentType: maybejson.Wrap(g.PaymentType),
		MaxPlayers:  g.MaxPlayers,
	}
	return json.Marshal(wg)
}
