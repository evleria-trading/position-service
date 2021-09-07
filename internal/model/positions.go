package model

import (
	"gopkg.in/guregu/null.v4"
	"time"
)

type Position struct {
	PositionID int64      `db:"position_id" json:"position_id"`
	AddPrice   float64    `db:"add_price" json:"add_price"`
	ClosePrice null.Float `db:"close_price" json:"close_price"`
	Symbol     string     `db:"symbol" json:"symbol"`
	OpenedAt   time.Time  `db:"opened_at" json:"-"`
	IsBuyType  bool       `db:"is_buy_type" json:"is_buy_type"`
	StopLoss   null.Float `db:"stop_loss" json:"stop_loss"`
	TakeProfit null.Float `db:"take_profit" json:"take_profit"`
	UserID     int64      `db:"user_id" json:"user_id"`
}

func (p *Position) IsClosed() bool {
	return p.ClosePrice.Valid
}

func (p *Position) GetFieldAddresses() []interface{} {
	return []interface{}{&p.PositionID, &p.AddPrice, &p.ClosePrice.NullFloat64, &p.Symbol, &p.OpenedAt, &p.IsBuyType, &p.StopLoss.NullFloat64, &p.TakeProfit.NullFloat64, &p.UserID}
}
