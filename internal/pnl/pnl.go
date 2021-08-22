package pnl

import (
	"context"
	"github.com/evleria/position-service/internal/model"
	"github.com/evleria/position-service/internal/pnl/profit"
	"github.com/evleria/position-service/internal/service"
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
	openedPositions := map[string]map[int64]channels{}
	for {
		select {
		case pr := <-pricesChan:
			for _, chans := range openedPositions[pr.Symbol] {
				chans.priceChan <- pr
			}
		case pos := <-openedChan:
			if _, ok := openedPositions[pos.Symbol]; !ok {
				openedPositions[pos.Symbol] = map[int64]channels{}
			}
			openedPositions[pos.Symbol][pos.PositionID] = m.handlePosition(pos)
		case pos := <-closedChan:
			if mp, ok := openedPositions[pos.Symbol]; ok {
				if chans, ok := mp[pos.PositionID]; ok {
					chans.closeChan <- struct{}{}
					delete(mp, pos.PositionID)
				}
			}
		case pos := <-updatedChan:
			if mp, ok := openedPositions[pos.Symbol]; ok {
				if chans, ok := mp[pos.PositionID]; ok {
					chans.updateChan <- pos
				}
			}
		}
	}
}

func (m *monitor) handlePosition(pos model.Position) channels {
	priceChan := make(chan model.Price)
	updateChan := make(chan model.Position)
	closeChan := make(chan struct{})
	go func() {
		for {
			select {
			case pr := <-priceChan:
				pnl := profit.Calculate(pos.AddPrice, pr.GetPrice(!pos.IsBuyType), pos.IsBuyType)
				log.WithFields(log.Fields{
					"positionId": pos.PositionID,
					"symbol":     pos.Symbol,
					"pnl":        pnl,
				}).Info("Calculated PnL")
				if (pos.StopLoss.Valid && pnl <= pos.StopLoss.Float64) || (pos.TakeProfit.Valid && pnl >= pos.TakeProfit.Float64) {
					_, err := m.positionService.ClosePosition(context.Background(), pos.PositionID, pr.Id)
					if err != nil {
						log.Error(err)
					}
				}
			case pos = <-updateChan:
			case <-closeChan:
				return
			}
		}
	}()
	return channels{
		priceChan:  priceChan,
		updateChan: updateChan,
		closeChan:  closeChan,
	}
}

type channels struct {
	priceChan  chan model.Price
	updateChan chan model.Position
	closeChan  chan struct{}
}
