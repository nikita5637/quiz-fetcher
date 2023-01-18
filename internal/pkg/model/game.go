package model

import (
	"time"
)

// Game ...
type Game struct {
	ExternalID  int32
	LeagueID    int32
	Type        int32
	Number      string
	Name        string
	PlaceID     int32
	DateTime    time.Time
	Price       uint32
	PaymentType string
	MaxPlayers  uint32
}
