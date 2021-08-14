package consumer

import (
	"context"
	"errors"
	"fmt"
	"github.com/evleria/PriceService/internal/model"
	"github.com/evleria/PriceService/internal/repository"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

type Price interface {
	Consume(ctx context.Context) error
}

type price struct {
	redis          *redis.Client
	repository     repository.Price
	warmupDuration time.Duration
}

func NewPriceConsumer(
	redisClient *redis.Client,
	priceRepository repository.Price,
	warmupDuration time.Duration) Price {
	return &price{
		redis:          redisClient,
		repository:     priceRepository,
		warmupDuration: warmupDuration,
	}
}

func (p *price) Consume(ctx context.Context) error {
	id := fmt.Sprintf("%d000-0", time.Now().Add(-p.warmupDuration).Unix())
	for {
		args := &redis.XReadArgs{
			Streams: []string{"prices", id},
		}
		r, err := p.redis.XRead(ctx, args).Result()
		if err != nil {
			return err
		}
		for _, message := range r[0].Messages {
			price, err := decodeMessage(message.Values)
			if err != nil {
				return err
			}
			p.repository.UpdatePrice(price)

			id = message.ID
		}
	}
}

func decodeMessage(values map[string]interface{}) (model.Price, error) {
	symbol, err := decodeString(values["symbol"])
	if err != nil {
		return model.Price{}, err
	}

	ask, err := decodeFloat64(values["ask"])
	if err != nil {
		return model.Price{}, err
	}

	bid, err := decodeFloat64(values["bid"])
	if err != nil {
		return model.Price{}, err
	}
	return model.Price{
		Symbol: symbol,
		Ask:    ask,
		Bid:    bid,
	}, nil
}

func decodeString(v interface{}) (string, error) {
	if v == nil {
		return "", errors.New("cannot decode nil")
	}
	if str, ok := v.(string); ok {
		return str, nil
	}
	return "", errors.New("cannot convert to string")
}

func decodeFloat64(v interface{}) (float64, error) {
	str, err := decodeString(v)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(str, 64)
}
