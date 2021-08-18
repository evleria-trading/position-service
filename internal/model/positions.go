package model

import (
	"database/sql"
	"time"
)

type Position struct {
	PositionID int64           `db:"position_id" json:"position_id"`
	AddPrice   float64         `db:"add_price" json:"add_price"`
	ClosePrice sql.NullFloat64 `db:"close_price" json:"-"`
	Symbol     string          `db:"symbol" json:"symbol"`
	OpenedAt   time.Time       `db:"opened_at" json:"-"`
	IsBuyType  bool            `db:"is_buy_type" json:"is_buy_type"`
}

func (p *Position) IsClosed() bool {
	return p.ClosePrice.Valid
}
