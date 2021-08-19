package pnl

import (
	"context"
	"github.com/evleria/position-service/internal/model"
	"github.com/evleria/position-service/internal/pnl/profit"
	"github.com/evleria/position-service/internal/service"
	log "github.com/sirupsen/logrus"
)

func CalculatePnlForOpenPositions(
	positionService service.Position,
	pricesChan chan model.Price,
	openedChan chan model.Position,
	closedChan chan model.Position,
	updatedChan chan model.Position) {
	openedPositions := map[string]map[int64]model.Position{}
	for {
		select {
		case pr := <-pricesChan:
			for id, pos := range openedPositions[pr.Symbol] {
				pnl := profit.Calculate(pos.AddPrice, pr.GetPrice(!pos.IsBuyType), pos.IsBuyType)
				log.WithFields(log.Fields{
					"positionId": id,
					"symbol":     pos.Symbol,
					"pnl":        pnl,
				}).Info("Calculated PnL")
				if (pos.StopLoss.Valid && pnl <= pos.StopLoss.Float64) || (pos.TakeProfit.Valid && pnl >= pos.TakeProfit.Float64) {
					_, err := positionService.ClosePosition(context.Background(), id, pr.Id)
					if err != nil {
						log.Error(err)
					}
				}
			}
		case pos := <-openedChan:
			if _, ok := openedPositions[pos.Symbol]; !ok {
				openedPositions[pos.Symbol] = map[int64]model.Position{}
			}
			openedPositions[pos.Symbol][pos.PositionID] = pos
		case pos := <-closedChan:
			delete(openedPositions[pos.Symbol], pos.PositionID)
		case pos := <-updatedChan:
			if m, ok := openedPositions[pos.Symbol]; ok {
				if _, ok := m[pos.PositionID]; ok {
					m[pos.PositionID] = pos
				}
			}
		}
	}
}
