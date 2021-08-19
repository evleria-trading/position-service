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

const (
	PositionOpenedChannel  = "notify_position_opened"
	PositionClosedChannel  = "notify_position_closed"
	PositionUpdatedChannel = "notify_position_updated"
)

type Position interface {
	CreatePosition(ctx context.Context, openPrice float64, symbol string, isBuyType bool) (int64, error)
	GetPositionByID(ctx context.Context, id int64) (*model.Position, error)
	ClosePosition(ctx context.Context, id int64, closePrice float64) error
	UpdateStopLoss(ctx context.Context, id int64, stopLoss float64) error
	UpdateTakeProfit(ctx context.Context, id int64, takeProfit float64) error
	ListenNotifications(ctx context.Context) (chan model.Position, chan model.Position, chan model.Position, error)
}

type position struct {
	db *pgxpool.Pool
}

func NewPositionRepository(db *pgxpool.Pool) Position {
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
		Scan(pos.GetFieldAddresses()...)
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

func (p *position) UpdateStopLoss(ctx context.Context, id int64, stopLoss float64) error {
	res, err := p.db.Exec(ctx, `UPDATE positions SET stop_loss=$1 WHERE position_id=$2 AND close_price IS NULL;`, stopLoss, id)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return ErrPositionNotFound
	}
	return nil
}

func (p *position) UpdateTakeProfit(ctx context.Context, id int64, takeProfit float64) error {
	res, err := p.db.Exec(ctx, `UPDATE positions SET take_profit=$1 WHERE position_id=$2 AND close_price IS NULL;`, takeProfit, id)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return ErrPositionNotFound
	}
	return nil
}

func (p *position) ListenNotifications(ctx context.Context) (openedChan chan model.Position, closedChan chan model.Position, updatedChan chan model.Position, err error) {
	conn, err := p.db.Acquire(ctx)
	if err != nil {
		return nil, nil, nil, err
	}

	err = listenChannels(ctx, conn, PositionOpenedChannel, PositionClosedChannel, PositionUpdatedChannel)
	if err != nil {
		return nil, nil, nil, err
	}

	openedChan = make(chan model.Position)
	closedChan = make(chan model.Position)
	updatedChan = make(chan model.Position)

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
				case PositionOpenedChannel:
					openedChan <- pos
				case PositionClosedChannel:
					closedChan <- pos
				case PositionUpdatedChannel:
					updatedChan <- pos
				}
			}
		}
	}()

	return openedChan, closedChan, updatedChan, nil
}

func listenChannels(ctx context.Context, conn *pgxpool.Conn, channels ...string) error {
	for _, channel := range channels {
		_, err := conn.Exec(ctx, "LISTEN "+channel+";")
		if err != nil {
			return err
		}
	}
	return nil
}

func decodePosition(payload string) (model.Position, error) {
	pos := model.Position{}
	err := json.Unmarshal([]byte(payload), &pos)
	return pos, err
}
