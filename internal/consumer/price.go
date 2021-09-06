package consumer

import (
	"context"
	"github.com/evleria-trading/position-service/internal/model"
	"github.com/evleria-trading/position-service/internal/repository"
	pricePb "github.com/evleria/price-service/protocol/pb"
	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
)

type Price interface {
	Consume(ctx context.Context) (chan model.Price, error)
}

type price struct {
	priceClient pricePb.PriceServiceClient
	repository  repository.Price
}

func NewPriceConsumer(priceClient pricePb.PriceServiceClient, priceRepository repository.Price) Price {
	return &price{
		priceClient: priceClient,
		repository:  priceRepository,
	}
}

func (p *price) Consume(ctx context.Context) (chan model.Price, error) {
	stream, err := p.priceClient.GetPrices(ctx, &empty.Empty{})
	if err != nil {
		return nil, err
	}
	ch := make(chan model.Price)
	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Error(err)
				return
			}
			pr := model.Price{
				Id:     msg.Id,
				Symbol: msg.Symbol,
				Ask:    msg.Ask,
				Bid:    msg.Bid,
			}

			log.WithFields(log.Fields{
				"id":     pr.Id,
				"symbol": pr.Symbol,
				"ask":    pr.Ask,
				"bid":    pr.Bid,
			}).Debug("Consumed price message")
			p.repository.UpdatePrice(pr)
			ch <- pr
		}
	}()

	return ch, nil
}
