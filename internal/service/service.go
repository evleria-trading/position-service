package service

import (
	"context"
	"github.com/evleria/PriceService/internal/repository"
)

type Position interface {
	AddPosition(ctx context.Context, price float64) (int, error)
}

type position struct {
	repository repository.Position
}

func NewPositionService(positionRepository repository.Position) Position {
	return &position{
		repository: positionRepository,
	}
}

func (p *position) AddPosition(ctx context.Context, price float64) (int, error) {
	id, err := p.repository.CreatePosition(ctx)
	return id, err
}
