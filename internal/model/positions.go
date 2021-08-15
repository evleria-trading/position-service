package model

import (
	"database/sql"
	"time"
)

type Position struct {
	PositionID int64           `db:"position_id"`
	AddPrice   float64         `db:"add_price"`
	ClosePrice sql.NullFloat64 `db:"close_price"`
	Symbol     string          `db:"symbol"`
	OpenedAt   time.Time       `db:"opened_at"`
	IsBuyType  bool            `db:"is_buy_type"`
}

func (p *Position) IsClosed() bool {
	return p.ClosePrice.Valid
}
