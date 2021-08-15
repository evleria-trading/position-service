package repository

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type Position interface {
	CreatePosition(ctx context.Context, openPrice float64, symbol string, isBuyType bool) (int, error)
}

type position struct {
	db *pgx.Conn
}

func NewPositionService(db *pgx.Conn) Position {
	return &position{
		db: db,
	}
}

func (p *position) CreatePosition(ctx context.Context, openPrice float64, symbol string, isBuyType bool) (int, error) {
	var id int
	err := p.db.QueryRow(
		ctx,
		`INSERT INTO positions (add_price, symbol, is_buy_type) VALUES ($1, $2, $3) RETURNING position_id;`,
		openPrice,
		symbol,
		isBuyType).Scan(&id)
	return id, err
}
