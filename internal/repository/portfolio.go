package repository

import (
	"errors"
	"github.com/evleria-trading/position-service/internal/model"
)

var (
	ErrPortfolioNotCalculated = errors.New("portfolio not calculated")
)

type Portfolio interface {
	UpdatePortfolio(position model.Position)
	RemoveFromPortfolio(position model.Position)
	RecalculatePortfolio(userId, positionId int64, price model.Price) (float64, error)

	GetPortfolioPnl(userId int64) (float64, error)
	GetAllPositions(userId int64) []model.Position
}

type portfolio struct {
	portfolios map[int64]*userPortfolio
}

func NewPortfolioRepository() Portfolio {
	return &portfolio{
		portfolios: map[int64]*userPortfolio{},
	}
}

func (p *portfolio) UpdatePortfolio(position model.Position) {
	up := p.getOrCreateUserPortfolio(position.UserID)
	up.AddPosition(position)
}

func (p *portfolio) RemoveFromPortfolio(position model.Position) {
	up := p.getOrCreateUserPortfolio(position.UserID)
	up.RemovePosition(position.PositionID)
}

func (p *portfolio) RecalculatePortfolio(userId, positionId int64, price model.Price) (float64, error) {
	up := p.getOrCreateUserPortfolio(userId)
	_ = up.UpdatePrice(positionId, price)
	return p.GetPortfolioPnl(userId)
}

func (p *portfolio) GetPortfolioPnl(userId int64) (float64, error) {
	up := p.getOrCreateUserPortfolio(userId)
	if up.NotCalculated() {
		return 0, ErrPortfolioNotCalculated
	}
	return up.pnlSum, nil
}

func (p *portfolio) GetAllPositions(userId int64) []model.Position {
	up := p.getOrCreateUserPortfolio(userId)
	return up.GetAllPositions()
}

func (p *portfolio) getOrCreateUserPortfolio(userId int64) *userPortfolio {
	if _, ok := p.portfolios[userId]; !ok {
		p.portfolios[userId] = newUserPortfolio()
	}
	return p.portfolios[userId]
}
