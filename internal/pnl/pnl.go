package pnl

import (
	"github.com/evleria/position-service/internal/model"
	log "github.com/sirupsen/logrus"
)

func CalculatePnlForOpenPositions(pricesChan chan model.Price, openedChan chan model.Position, closedChan chan model.Position) {
	openedPositions := map[string]map[int64]model.Position{}
	for {
		select {
		case pr := <-pricesChan:
			for id, pos := range openedPositions[pr.Symbol] {
				log.WithFields(log.Fields{
					"positionId": id,
					"symbol":     pos.Symbol,
					"pnl":        Calculate(pos.AddPrice, pr.GetPrice(!pos.IsBuyType), pos.IsBuyType),
				}).Info("Calculated PnL")
			}
		case pos := <-openedChan:
			if _, ok := openedPositions[pos.Symbol]; !ok {
				openedPositions[pos.Symbol] = map[int64]model.Position{}
			}
			openedPositions[pos.Symbol][pos.PositionID] = pos
		case pos := <-closedChan:
			delete(openedPositions[pos.Symbol], pos.PositionID)
		}
	}
}

func Calculate(openPrice, closePrice float64, isBuyType bool) float64 {
	if isBuyType {
		return closePrice - openPrice
	}
	return openPrice - closePrice
}
