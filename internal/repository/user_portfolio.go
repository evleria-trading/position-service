package repository

import (
	"github.com/evleria-trading/position-service/internal/model"
	"github.com/evleria-trading/position-service/internal/pnl/profit"
)

type userPortfolio struct {
	positions              map[int64]*positionWithPnl
	pnlSum                 float64
	notCalculatedPositions int
}

func newUserPortfolio() *userPortfolio {
	return &userPortfolio{
		positions: map[int64]*positionWithPnl{},
	}
}

func (up *userPortfolio) AddPosition(position model.Position) {
	up.positions[position.PositionID] = &positionWithPnl{Position: position}
	up.notCalculatedPositions++
}

func (up *userPortfolio) RemovePosition(positionId int64) {
	if p, ok := up.positions[positionId]; ok {
		if !p.pnlCalculated {
			up.notCalculatedPositions--
		}

		up.pnlSum -= p.pnl
		delete(up.positions, positionId)
	}
}

func (up *userPortfolio) UpdatePrice(positionId int64, price model.Price) error {
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
		return nil
	} else {
		return ErrPositionNotFound
	}
}

func (up *userPortfolio) NotCalculated() bool {
	return up.notCalculatedPositions > 0
}

func (up *userPortfolio) GetAllPositions() []model.Position {
	result := make([]model.Position, 0, len(up.positions))
	for _, pos := range up.positions {
		result = append(result, pos.Position)
	}
	return result
}

type positionWithPnl struct {
	model.Position
	pnl           float64
	pnlCalculated bool
}
