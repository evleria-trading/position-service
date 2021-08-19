package profit

func Calculate(openPrice, closePrice float64, isBuyType bool) float64 {
	if isBuyType {
		return closePrice - openPrice
	}
	return openPrice - closePrice
}
