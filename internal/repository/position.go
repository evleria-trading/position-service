package repository

import (
	"context"
	"errors"
	"github.com/evleria/PriceService/internal/model"
	"github.com/jackc/pgx/v4"
)

var (
	ErrPositionNotFound      = errors.New("position not found")
	ErrPositionAlreadyClosed = errors.New("position already closed")
)

type Position interface {
	CreatePosition(ctx context.Context, openPrice float64, symbol string, isBuyType bool) (int64, error)
	GetPositionByID(ctx context.Context, id int64) (*model.Position, error)
	ClosePosition(ctx context.Context, id int64, closePrice float64) error
}

type position struct {
	db *pgx.Conn
}

func NewPositionService(db *pgx.Conn) Position {
	return &position{
		db: db,
	}
}

func (p *position) CreatePosition(ctx context.Context, openPrice float64, symbol string, isBuyType bool) (int64, error) {
	var id int64
	err := p.db.QueryRow(
		ctx,
		`INSERT INTO positions (add_price, symbol, is_buy_type) VALUES ($1, $2, $3) RETURNING position_id;`,
		openPrice,
		symbol,
		isBuyType).Scan(&id)
	return id, err
}

func (p *position) GetPositionByID(ctx context.Context, id int64) (*model.Position, error) {
	pos := model.Position{}
	err := p.db.QueryRow(ctx, `SELECT * FROM positions WHERE position_id=$1;`, id).
		Scan(&pos.PositionID, &pos.AddPrice, &pos.ClosePrice, &pos.Symbol, &pos.OpenedAt, &pos.IsBuyType)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, ErrPositionNotFound
		}
		return nil, err
	}

	return &pos, nil
}

func (p *position) ClosePosition(ctx context.Context, id int64, closePrice float64) error {
	res, err := p.db.Exec(ctx, `UPDATE positions SET close_price=$1 WHERE position_id=$2 AND close_price IS NULL;`, closePrice, id)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return ErrPositionAlreadyClosed
	}
	return nil
}
