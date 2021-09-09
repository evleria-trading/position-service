package repository

import (
	"errors"
	"github.com/evleria-trading/position-service/internal/model"
	"github.com/evleria-trading/position-service/internal/pnl/profit"
)

var (
	ErrPortfolioNotCalculated = errors.New("portfolio not calculated")
)

type Portfolio interface {
	UpdatePortfolio(position model.Position, price model.Price)
	UpdatePortfolioWithoutPrice(position model.Position)
	RemoveFromPortfolio(position model.Position)
	RecalculatePortfolio(userId, positionId int64, price model.Price)

	GetPortfolioBalance(userId int64) (float64, error)
}

type portfolio struct {
	portfolios map[int64]*userPortfolio
}

func NewPortfolioRepository() Portfolio {
	return &portfolio{
		portfolios: map[int64]*userPortfolio{},
	}
}

func (p *portfolio) UpdatePortfolio(position model.Position, price model.Price) {
	up := p.getOrCreateUserPortfolio(position.UserID)
	up.AddPosition(position)
	up.UpdatePrice(position.PositionID, price)
}

func (p *portfolio) UpdatePortfolioWithoutPrice(position model.Position) {
	up := p.getOrCreateUserPortfolio(position.UserID)
	up.AddPosition(position)
}

func (p *portfolio) RemoveFromPortfolio(position model.Position) {
	up := p.getOrCreateUserPortfolio(position.UserID)
	up.RemovePosition(position)
}

func (p *portfolio) RecalculatePortfolio(userId, positionId int64, price model.Price) {
	up := p.getOrCreateUserPortfolio(userId)
	up.UpdatePrice(positionId, price)
}

func (p *portfolio) GetPortfolioBalance(userId int64) (float64, error) {
	up := p.getOrCreateUserPortfolio(userId)
	if up.NotCalculated() {
		return 0, ErrPortfolioNotCalculated
	}
	return up.pnlSum, nil
}

func (p *portfolio) getOrCreateUserPortfolio(userId int64) *userPortfolio {
	if _, ok := p.portfolios[userId]; !ok {
		p.portfolios[userId] = newUserPortfolio()
	}
	return p.portfolios[userId]
}

type userPortfolio struct {
	positions              map[int64]positionWithPnl
	pnlSum                 float64
	notCalculatedPositions int
}

func newUserPortfolio() *userPortfolio {
	return &userPortfolio{
		positions: map[int64]positionWithPnl{},
	}
}

func (up *userPortfolio) AddPosition(position model.Position) {
	up.positions[position.PositionID] = positionWithPnl{Position: position}
	up.notCalculatedPositions++
}

func (up *userPortfolio) RemovePosition(position model.Position) {
	if p, ok := up.positions[position.PositionID]; ok {
		if !p.pnlCalculated {
			up.notCalculatedPositions--
		}

		up.pnlSum -= p.pnl
		delete(up.positions, position.PositionID)
	}
}

func (up *userPortfolio) UpdatePrice(positionId int64, price model.Price) {
	if pos, ok := up.positions[positionId]; ok {
		newPnl := profit.Calculate(pos.AddPrice, price.GetPrice(!pos.IsBuyType), pos.IsBuyType)

		if pos.pnlCalculated {
			up.pnlSum += newPnl - pos.pnl
		} else {
			up.pnlSum += newPnl
			up.notCalculatedPositions--
			pos.pnlCalculated = true
		}

		pos.pnl = newPnl
	}
}

func (up *userPortfolio) NotCalculated() bool {
	return up.notCalculatedPositions > 0
}

type positionWithPnl struct {
	model.Position
	pnl           float64
	pnlCalculated bool
}
