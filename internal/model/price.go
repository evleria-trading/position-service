package model

type Price struct {
	Id     string
	Ask    float64
	Bid    float64
	Symbol string
}

func (p *Price) GetPrice(isBuy bool) float64 {
	if isBuy {
		return p.Ask
	}
	return p.Bid
}
