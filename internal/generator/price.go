package generator

import (
	"context"
	"github.com/evleria/PriceService/internal/model"
	"github.com/evleria/PriceService/internal/producer"
	"math"
	"math/rand"
	"time"
)

type Price interface {
	GeneratePrices(ctx context.Context) error
}

type price struct {
	pricesProducer producer.Price
}

func NewPricesGenerator(pricesProducer producer.Price) Price {
	return &price{
		pricesProducer: pricesProducer,
	}
}

var symbols = []string{
	"AAPL",
	"TSLA",
	"GOOG",
	"MSFT",
	"AMZN",
	"FB",
	"NVDA",
}

func (g *price) GeneratePrices(ctx context.Context) error {
	m := map[string]model.Price{}
	for {
		select {
		case <-time.After(time.Millisecond * 250):
			s := symbols[rand.Intn(len(symbols))]
			var p model.Price
			if prev, ok := m[s]; ok {
				c := getRandomCoefficient()
				p = newPrice(s, prev.Ask*c, prev.Bid*c)
			} else {
				base := 200 + rand.Float64()*500
				p = newPrice(s, base, base*0.9)
			}
			m[s] = p

			err := g.pricesProducer.Produce(ctx, p)
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return nil
		}
	}
}

func newPrice(symbol string, ask, bid float64) model.Price {
	return model.Price{
		Symbol: symbol,
		Ask:    roundToCents(ask),
		Bid:    roundToCents(bid),
	}
}

func roundToCents(f float64) float64 {
	return math.Round(f*100) / 100
}

func getRandomCoefficient() float64 {
	return rand.Float64()*0.2 + 0.9
}
