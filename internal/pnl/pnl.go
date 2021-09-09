package pnl

import (
	"github.com/evleria-trading/position-service/internal/model"
	"github.com/evleria-trading/position-service/internal/service"
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

	openedPositions := map[string]map[int64]*positionActor{}

	for {
		select {
		case pr := <-pricesChan:
			for _, pa := range openedPositions[pr.Symbol] {
				pa.PriceChanged(pr)
				m.positionService.RecalculatePortfolio(pa.userId, pa.positionId, pr)
			}

		case pos := <-openedChan:
			if _, ok := openedPositions[pos.Symbol]; !ok {
				openedPositions[pos.Symbol] = map[int64]*positionActor{}
			}

			positionActor := NewPositionActor(pos, m.positionService)
			openedPositions[pos.Symbol][pos.PositionID] = positionActor

			m.positionService.UpdatePortfolio(pos)

		case pos := <-closedChan:
			if mp, ok := openedPositions[pos.Symbol]; ok {
				if pa, ok := mp[pos.PositionID]; ok {
					pa.Close()
					delete(mp, pos.PositionID)
					m.positionService.RemoveFromPortfolio(pos)
				}
			}

		case pos := <-updatedChan:
			if mp, ok := openedPositions[pos.Symbol]; ok {
				if pa, ok := mp[pos.PositionID]; ok {
					pa.Update(pos)
				}
			}
		}
	}
}
