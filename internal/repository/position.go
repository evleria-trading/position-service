package repository

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type Position interface {
	CreatePosition(ctx context.Context) (int, error)
}

type position struct {
	db *pgx.Conn
}

func NewPositionService(db *pgx.Conn) Position {
	return &position{
		db: db,
	}
}

func (p *position) CreatePosition(ctx context.Context) (int, error) {
	var id int
	err := p.db.QueryRow(ctx, `INSERT INTO positions DEFAULT VALUES RETURNING position_id`).Scan(&id)
	return id, err
}
