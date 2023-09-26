package model

import (
	"time"

	"github.com/mono83/maybe"
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
