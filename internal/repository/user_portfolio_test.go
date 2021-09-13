package repository

import (
	"github.com/evleria-trading/position-service/internal/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdatingPortfolio(t *testing.T) {
	up := newUserPortfolio()
	up.AddPosition(model.Position{
		PositionID: 1,
		IsBuyType:  true,
		AddPrice:   50.0,
	})

	assert.True(t, up.NotCalculated())

	err := up.UpdatePrice(1, model.Price{
		Bid: 45,
		Ask: 55,
	})

	assert.NoError(t, err)
	assert.False(t, up.NotCalculated())
	assert.Equal(t, -5., up.pnlSum)

	err = up.UpdatePrice(1, model.Price{
		Bid: 46,
		Ask: 56,
	})

	assert.NoError(t, err)
	assert.False(t, up.NotCalculated())
	assert.Equal(t, -4., up.pnlSum)

	up.RemovePosition(1)

	assert.False(t, up.NotCalculated())
	assert.Equal(t, 0., up.pnlSum)
}
