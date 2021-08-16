package service

import (
	"context"
	"errors"
	"github.com/evleria/PriceService/internal/model"
	"github.com/evleria/PriceService/internal/repository"
	log "github.com/sirupsen/logrus"
)

var (
	ErrPriceNotFound    = errors.New("price not found")
	ErrPositionNotFound = errors.New("position not found")
	ErrPositionClosed   = errors.New("position closed")
)

type Position interface {
	AddPosition(ctx context.Context, symbol string, isBuyType bool) (int64, error)
	ClosePosition(ctx context.Context, id int64) (float64, error)
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

func (p *position) AddPosition(ctx context.Context, symbol string, isBuyType bool) (int64, error) {
	price, err := p.priceRepository.GetPrice(symbol)
	if err != nil {
		return 0, ErrPriceNotFound
	}

	openPrice := getPrice(price, isBuyType)
	id, err := p.positionRepository.CreatePosition(ctx, openPrice, symbol, isBuyType)
	if err != nil {
		return 0, err
	}

	log.WithFields(log.Fields{"id": id}).Info("Created position")
	return id, nil
}

func (p *position) ClosePosition(ctx context.Context, id int64) (float64, error) {
	pos, err := p.positionRepository.GetPositionByID(ctx, id)
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

	closePrice := getPrice(price, !pos.IsBuyType)
	err = p.positionRepository.ClosePosition(ctx, id, closePrice)
	if err != nil {
		if err == repository.ErrPositionAlreadyClosed {
			return 0, ErrPositionClosed
		}
		return 0, err
	}

	log.WithFields(log.Fields{"id": id}).Info("Closed position")
	if pos.IsBuyType {
		return closePrice - pos.AddPrice, nil
	}
	return pos.AddPrice - closePrice, nil
}

func getPrice(price model.Price, isBuy bool) float64 {
	if isBuy {
		return price.Ask
	}
	return price.Bid
}
