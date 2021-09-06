package service

import (
	"context"
	"errors"
	"github.com/evleria-trading/position-service/internal/model"
	"github.com/evleria-trading/position-service/internal/pnl/profit"
	"github.com/evleria-trading/position-service/internal/repository"
	log "github.com/sirupsen/logrus"
)

var (
	ErrPriceNotFound           = errors.New("price not found")
	ErrPositionNotFound        = errors.New("position not found")
	ErrPositionClosed          = errors.New("position closed")
	ErrPriceChanged            = errors.New("price is changed")
	ErrStopLossIsNotNegative   = errors.New("stop loss is not negative")
	ErrTakeProfitIsNotPositive = errors.New("take profit is not positive")
)

type Position interface {
	AddPosition(ctx context.Context, symbol string, isBuyType bool, priceId string) (int64, error)
	ClosePosition(ctx context.Context, positionId int64, priceId string) (float64, error)
	GetOpenPosition(ctx context.Context, positionId int64) (*model.Position, error)
	SetStopLoss(ctx context.Context, positionId int64, stopLoss float64) error
	SetTakeProfit(ctx context.Context, id int64, takeProfit float64) error
}

type position struct {
	positionRepository repository.Position
	priceRepository    repository.Price
}

func NewPositionService(positionRepository repository.Position, priceRepository repository.Price) Position {
	return &position{
		positionRepository: positionRepository,
		priceRepository:    priceRepository,
	}
}

func (p *position) AddPosition(ctx context.Context, symbol string, isBuyType bool, priceId string) (int64, error) {
	price, err := p.priceRepository.GetPrice(symbol)
	if err != nil {
		return 0, ErrPriceNotFound
	}

	if price.Id != priceId {
		return 0, ErrPriceChanged
	}

	openPrice := price.GetPrice(isBuyType)
	id, err := p.positionRepository.CreatePosition(ctx, openPrice, symbol, isBuyType)
	if err != nil {
		return 0, err
	}

	log.WithFields(log.Fields{"id": id}).Info("Created position")
	return id, nil
}

func (p *position) ClosePosition(ctx context.Context, positionId int64, priceId string) (float64, error) {
	pos, err := p.positionRepository.GetPositionByID(ctx, positionId)
	if err != nil {
		if err == repository.ErrPositionNotFound {
			return 0, ErrPositionNotFound
		}
		return 0, err
	}

	if pos.IsClosed() {
		return 0, ErrPositionClosed
	}

	price, err := p.priceRepository.GetPrice(pos.Symbol)
	if err != nil {
		return 0, ErrPriceNotFound
	}

	if price.Id != priceId {
		return 0, ErrPriceChanged
	}

	closePrice := price.GetPrice(!pos.IsBuyType)
	err = p.positionRepository.ClosePosition(ctx, positionId, closePrice)
	if err != nil {
		if err == repository.ErrPositionAlreadyClosed {
			return 0, ErrPositionClosed
		}
		return 0, err
	}

	log.WithFields(log.Fields{"id": positionId}).Info("Closed position")
	return profit.Calculate(pos.AddPrice, closePrice, pos.IsBuyType), nil
}

func (p *position) GetOpenPosition(ctx context.Context, positionId int64) (*model.Position, error) {
	pos, err := p.positionRepository.GetPositionByID(ctx, positionId)
	if err != nil {
		if err == repository.ErrPositionNotFound {
			return nil, ErrPositionNotFound
		}
		return nil, err
	}

	if pos.IsClosed() {
		return nil, ErrPositionClosed
	}
	return pos, nil
}

func (p *position) SetStopLoss(ctx context.Context, positionId int64, stopLoss float64) error {
	if stopLoss >= 0 {
		return ErrStopLossIsNotNegative
	}
	err := p.positionRepository.UpdateStopLoss(ctx, positionId, stopLoss)
	if err == repository.ErrPositionNotFound {
		return ErrPositionNotFound
	}
	return err
}

func (p *position) SetTakeProfit(ctx context.Context, positionIid int64, takeProfit float64) error {
	if takeProfit <= 0 {
		return ErrTakeProfitIsNotPositive
	}
	err := p.positionRepository.UpdateTakeProfit(ctx, positionIid, takeProfit)
	if err == repository.ErrPositionNotFound {
		return ErrPositionNotFound
	}
	return err
}
