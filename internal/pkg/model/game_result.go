package model

import "github.com/mono83/maybe"

// GameResult ...
type GameResult struct {
	ResultPlace uint32
	RoundPoints maybe.Maybe[string]
}
