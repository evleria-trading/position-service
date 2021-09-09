package pnl

import (
	"context"
	"github.com/evleria-trading/position-service/internal/model"
	"github.com/evleria-trading/position-service/internal/pnl/profit"
	"github.com/evleria-trading/position-service/internal/service"
	log "github.com/sirupsen/logrus"
)

type positionActor struct {
	positionId int64
	userId     int64

	priceChan  chan model.Price
	updateChan chan model.Position
	closeChan  chan struct{}
}

func NewPositionActor(pos model.Position, positionService service.Position) *positionActor {
	priceChan := make(chan model.Price)
	updateChan := make(chan model.Position)
	closeChan := make(chan struct{})
	go func() {
		for {
			//lastPnl := math.MaxFloat64
			select {
			case pr := <-priceChan:
				pnl := profit.Calculate(pos.AddPrice, pr.GetPrice(!pos.IsBuyType), pos.IsBuyType)
				if (pos.StopLoss.Valid && pnl <= pos.StopLoss.Float64) || (pos.TakeProfit.Valid && pnl >= pos.TakeProfit.Float64) {
					_, err := positionService.ClosePosition(context.Background(), pos.UserID, pos.PositionID, pr.Id)
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

	return &positionActor{
		positionId: pos.PositionID,
		userId:     pos.UserID,

		priceChan:  priceChan,
		updateChan: updateChan,
		closeChan:  closeChan,
	}
}

func (p *positionActor) PriceChanged(price model.Price) {
	p.priceChan <- price
}

func (p *positionActor) Update(position model.Position) {
	p.updateChan <- position
}

func (p *positionActor) Close() {
	p.closeChan <- struct{}{}
}
