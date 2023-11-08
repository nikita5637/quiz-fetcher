package info

import "time"

// StartedAt ...
var StartedAt time.Time

func init() {
	StartedAt = time.Now().UTC()
}
