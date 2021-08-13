package service

import (
	"context"
	"github.com/evleria/PriceService/internal/producer"
	"github.com/evleria/PriceService/internal/repository"
)

type Position interface {
	AddPosition(ctx context.Context, price float64) (int, error)
}

type position struct {
	repository    repository.Position
	priceProducer producer.Price
}

func NewPositionService(positionRepository repository.Position, priceProducer producer.Price) Position {
	return &position{
		repository:    positionRepository,
		priceProducer: priceProducer,
	}
}

func (p *position) AddPosition(ctx context.Context, price float64) (int, error) {
	id, err := p.repository.CreatePosition(ctx)
	return id, err
}
