package producer

import (
	"context"
	"fmt"
	"github.com/evleria/PriceService/internal/model"
	"github.com/go-redis/redis/v8"
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
	fmt.Println("producing message", price)
	args := &redis.XAddArgs{
		Stream: "prices",
		Values: map[string]interface{}{
			"ask":    price.Ask,
			"bid":    price.Bid,
			"symbol": price.Symbol,
		},
	}
	return p.redis.XAdd(ctx, args).Err()
}
