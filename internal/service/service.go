package service

import (
	"context"
	"errors"
	"github.com/evleria/position-service/internal/model"
	"github.com/evleria/position-service/internal/pnl"
	"github.com/evleria/position-service/internal/repository"
	log "github.com/sirupsen/logrus"
)

var (
	ErrPriceNotFound    = errors.New("price not found")
	ErrPositionNotFound = errors.New("position not found")
	ErrPositionClosed   = errors.New("position closed")
	ErrPriceChanged     = errors.New("price is changed")
)

type Position interface {
	AddPosition(ctx context.Context, symbol string, isBuyType bool, priceId string) (int64, error)
	ClosePosition(ctx context.Context, positionId int64, priceId string) (float64, error)
	GetOpenPosition(ctx context.Context, positionId int64) (*model.Position, error)
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
	return pnl.Calculate(pos.AddPrice, closePrice, pos.IsBuyType), nil
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
