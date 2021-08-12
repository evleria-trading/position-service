package entities

import "time"

type Position struct {
	PositionID   int       `db:"position_id"`
	AddPriceID   string    `db:"add_price_id"`
	ClosePriceID string    `db:"close_price_id"`
	OpenedAt     time.Time `db:"opened_at"`
}
