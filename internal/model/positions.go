package model

import "time"

type Position struct {
	PositionID int       `db:"position_id"`
	AddPrice   float64   `db:"add_price"`
	ClosePrice float64   `db:"close_price"`
	Symbol     string    `db:"symbol"`
	OpenedAt   time.Time `db:"opened_at"`
}
