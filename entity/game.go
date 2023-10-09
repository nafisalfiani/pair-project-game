package entity

import (
	"time"
)

// Game represents the "game" table
type Game struct {
	GameID        int
	Name          string
	Description   string
	PublishedDate time.Time
	Rating        float64
}
