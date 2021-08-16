package producer

import (
	"context"
	"github.com/evleria/position-service/internal/model"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

type Price interface {
	Produce(ctx context.Context, price model.Price) error
}

type price struct {
	redis *redis.Client
}

func NewProducerPrice(redisClient *redis.Client) Price {
	return &price{
		redis: redisClient,
	}
}

func (p *price) Produce(ctx context.Context, price model.Price) error {
	values := map[string]interface{}{
		"ask":    price.Ask,
		"bid":    price.Bid,
		"symbol": price.Symbol,
	}
	args := &redis.XAddArgs{
		Stream: "prices",
		Values: values,
	}
	err := p.redis.XAdd(ctx, args).Err()
	if err != nil {
		return err
	}
	log.WithFields(values).Debug("Produced price message")
	return nil
}
