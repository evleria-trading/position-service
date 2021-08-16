package repository

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/evleria/position-service/internal/model"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
)

var (
	ErrPositionNotFound      = errors.New("position not found")
	ErrPositionAlreadyClosed = errors.New("position already closed")
)

type Position interface {
	CreatePosition(ctx context.Context, openPrice float64, symbol string, isBuyType bool) (int64, error)
	GetPositionByID(ctx context.Context, id int64) (*model.Position, error)
	ClosePosition(ctx context.Context, id int64, closePrice float64) error
	ListenNotifications(ctx context.Context) (chan model.Position, chan model.Position, error)
}

type position struct {
	db *pgxpool.Pool
}

func NewPositionRepository(db *pgxpool.Pool) Position {
	res := &position{
		db: db,
	}
	return res
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

func (p *position) ListenNotifications(ctx context.Context) (openedChan chan model.Position, closedChan chan model.Position, err error) {
	conn, err := p.db.Acquire(ctx)
	if err != nil {
		return nil, nil, err
	}
	_, err = conn.Exec(ctx, "LISTEN notify_position_opened;")
	if err != nil {
		return nil, nil, err
	}
	_, err = conn.Exec(ctx, "LISTEN notify_position_closed;")
	if err != nil {
		return nil, nil, err
	}

	openedChan = make(chan model.Position)
	closedChan = make(chan model.Position)

	go func() {
		for {
			notification, err := conn.Conn().WaitForNotification(ctx)
			if err != nil {
				log.Error(err)
			} else {
				pos, err := decodePosition(notification.Payload)
				if err != nil {
					log.Error(err)
				}
				switch notification.Channel {
				case "notify_position_opened":
					openedChan <- pos
				case "notify_position_closed":
					closedChan <- pos
				}
			}
		}
	}()

	return openedChan, closedChan, nil
}

func decodePosition(payload string) (model.Position, error) {
	pos := model.Position{}
	err := json.Unmarshal([]byte(payload), &pos)
	return pos, err
}
