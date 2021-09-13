package pnl

import (
	"context"
	"github.com/evleria-trading/position-service/internal/model"
	"github.com/evleria-trading/position-service/internal/pnl/profit"
	"github.com/evleria-trading/position-service/internal/service"
	log "github.com/sirupsen/logrus"
)

type Monitor interface {
	CalculatePnlForOpenPositions(
		pricesChan chan model.Price,
		openedChan chan model.Position,
		closedChan chan model.Position,
		updatedChan chan model.Position)
}

type monitor struct {
	positionService service.Position
}

func NewMonitor(positionService service.Position) Monitor {
	return &monitor{
		positionService: positionService,
	}
}

func (m *monitor) CalculatePnlForOpenPositions(
	pricesChan chan model.Price,
	openedChan chan model.Position,
	closedChan chan model.Position,
	updatedChan chan model.Position) {

	openedPositions := map[string]map[int64]model.Position{}

	for {
		select {
		case pr := <-pricesChan:
			for _, pos := range openedPositions[pr.Symbol] { // read outer, read inner
				pnl := profit.Calculate(pos.AddPrice, pr.GetPrice(!pos.IsBuyType), pos.IsBuyType)
				if (pos.StopLoss.Valid && pnl <= pos.StopLoss.Float64) || (pos.TakeProfit.Valid && pnl >= pos.TakeProfit.Float64) {
					_, err := m.positionService.ClosePosition(context.Background(), pos.UserID, pos.PositionID, pr.Id)
					if err != nil {
						log.Warn(err)
					}
				} else {
					err := m.positionService.RecalculatePortfolio(context.Background(), pos.UserID, pos.PositionID, pr)
					if err != nil {
						log.Warn(err)
					}
				}
			}

		case pos := <-openedChan:
			if _, ok := openedPositions[pos.Symbol]; !ok { // read outer
				openedPositions[pos.Symbol] = map[int64]model.Position{} // write outer
			}
			openedPositions[pos.Symbol][pos.PositionID] = pos // write inner
			err := m.positionService.UpdatePortfolio(context.Background(), pos)
			if err != nil {
				log.Warn(err)
			}

		case pos := <-closedChan:
			if mp, ok := openedPositions[pos.Symbol]; ok { // read outer
				if _, ok := mp[pos.PositionID]; ok { // read inner
					delete(mp, pos.PositionID) // write inner
					m.positionService.RemoveFromPortfolio(pos)
				}
			}

		case pos := <-updatedChan:
			if mp, ok := openedPositions[pos.Symbol]; ok { // read outer
				if _, ok := mp[pos.PositionID]; ok { // read inner
					mp[pos.PositionID] = pos // write inner
				}
			}
		}
	}
}
