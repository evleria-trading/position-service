package entities

import "time"

type Position struct {
	PositionID   int       `db:"id"`
	AddPriceID   string    `db:"add_price_id"`
	ClosePriceID string    `db:"close_price_id"`
	OpenAt       time.Time `db:"open_at"`
}
