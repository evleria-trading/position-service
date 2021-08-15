package service

import (
	"context"
	"errors"
	"github.com/evleria/PriceService/internal/model"
	"github.com/evleria/PriceService/internal/repository"
)

var ErrPriceNotFound = errors.New("price not found")

type Position interface {
	AddPosition(ctx context.Context, symbol string, isBuyType bool) (int, error)
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

func (p *position) AddPosition(ctx context.Context, symbol string, isBuyType bool) (int, error) {
	price, err := p.priceRepository.GetPrice(symbol)
	if err != nil {
		return 0, ErrPriceNotFound
	}

	openPrice := getPrice(price, isBuyType)
	return p.positionRepository.CreatePosition(ctx, openPrice, symbol, isBuyType)
}

func getPrice(price model.Price, isBuyType bool) float64 {
	if isBuyType {
		return price.Ask
	}
	return price.Bid
}
