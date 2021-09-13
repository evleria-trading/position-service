package service

import (
	"context"
	"errors"
	"github.com/evleria-trading/position-service/internal/model"
	"github.com/evleria-trading/position-service/internal/pnl/profit"
	"github.com/evleria-trading/position-service/internal/repository"
	log "github.com/sirupsen/logrus"
)

var (
	ErrPriceNotFound           = errors.New("price not found")
	ErrPositionNotFound        = errors.New("position not found")
	ErrPositionClosed          = errors.New("position closed")
	ErrPriceChanged            = errors.New("price is changed")
	ErrStopLossIsNotNegative   = errors.New("stop loss is not negative")
	ErrTakeProfitIsNotPositive = errors.New("take profit is not positive")
	ErrInsufficientBalance     = errors.New("insufficient balance")
)

type Position interface {
	AddPosition(ctx context.Context, userId int64, symbol string, isBuyType bool, priceId string) (int64, error)
	ClosePosition(ctx context.Context, userId int64, positionId int64, priceId string) (float64, error)
	GetOpenPosition(ctx context.Context, positionId int64) (*model.Position, error)
	SetStopLoss(ctx context.Context, userId int64, positionId int64, stopLoss float64) error
	SetTakeProfit(ctx context.Context, userId int64, positionId int64, takeProfit float64) error

	UpdatePortfolio(ctx context.Context, position model.Position) error
	RemoveFromPortfolio(position model.Position)
	RecalculatePortfolio(ctx context.Context, userId, positionId int64, price model.Price) error
}

type position struct {
	positionRepository  repository.Position
	priceRepository     repository.Price
	userRepository      repository.User
	portfolioRepository repository.Portfolio
}

func NewPositionService(
	positionRepository repository.Position,
	priceRepository repository.Price,
	userRepository repository.User,
	portfolioRepository repository.Portfolio) Position {
	return &position{
		positionRepository:  positionRepository,
		priceRepository:     priceRepository,
		userRepository:      userRepository,
		portfolioRepository: portfolioRepository,
	}
}

func (p *position) AddPosition(ctx context.Context, userId int64, symbol string, isBuyType bool, priceId string) (int64, error) {
	price, err := p.priceRepository.GetPrice(symbol)
	if err != nil {
		return 0, ErrPriceNotFound
	}

	if price.Id != priceId {
		return 0, ErrPriceChanged
	}

	openPrice := price.GetPrice(isBuyType)
	if isBuyType {
		can, err := p.canAfford(ctx, userId, openPrice)
		if err != nil {
			return 0, err
		}
		if !can {
			return 0, ErrInsufficientBalance
		}
	}

	id, err := p.positionRepository.CreatePosition(ctx, userId, openPrice, symbol, isBuyType)
	if err != nil {
		return 0, err
	}

	balanceChange := openPrice
	if isBuyType {
		balanceChange = -openPrice
	}

	_, err = p.userRepository.AddToBalance(ctx, userId, balanceChange)
	if err != nil {
		return 0, err
	}

	log.WithFields(log.Fields{"id": id}).Info("Created position")
	return id, nil
}

func (p *position) canAfford(ctx context.Context, userId int64, price float64) (bool, error) {
	bal, err := p.userRepository.GetBalance(ctx, userId)
	if err != nil {
		return false, err
	}
	pnl, err := p.portfolioRepository.GetPortfolioPnl(userId)
	if err != nil {
		return false, err
	}
	return bal+pnl >= price, nil
}

func (p *position) ClosePosition(ctx context.Context, userId int64, positionId int64, priceId string) (float64, error) {
	pos, err := p.positionRepository.GetPositionByID(ctx, positionId)
	if err != nil {
		if err == repository.ErrPositionNotFound {
			return 0, ErrPositionNotFound
		}
		return 0, err
	}

	if pos.UserID != userId {
		return 0, ErrPositionNotFound
	}

	if pos.IsClosed() {
		return 0, ErrPositionClosed
	}

	price, err := p.priceRepository.GetPrice(pos.Symbol)
	if err != nil {
		return 0, ErrPriceNotFound
	}

	if price.Id != priceId {
		return 0, ErrPriceChanged
	}

	closePrice := price.GetPrice(!pos.IsBuyType)
	err = p.positionRepository.ClosePosition(ctx, positionId, closePrice)
	if err != nil {
		if err == repository.ErrPositionAlreadyClosed {
			return 0, ErrPositionClosed
		}
		return 0, err
	}

	log.WithFields(log.Fields{"id": positionId}).Info("Closed position")

	balanceChange := p.getClosingBalanceChange(pos.IsBuyType, closePrice)
	_, err = p.userRepository.AddToBalance(ctx, userId, balanceChange)

	closeProfit := profit.Calculate(pos.AddPrice, closePrice, pos.IsBuyType)
	return closeProfit, err
}

func (p *position) GetOpenPosition(ctx context.Context, positionId int64) (*model.Position, error) {
	pos, err := p.positionRepository.GetPositionByID(ctx, positionId)
	if err != nil {
		if err == repository.ErrPositionNotFound {
			return nil, ErrPositionNotFound
		}
		return nil, err
	}

	if pos.IsClosed() {
		return nil, ErrPositionClosed
	}
	return pos, nil
}

func (p *position) SetStopLoss(ctx context.Context, userId int64, positionId int64, stopLoss float64) error {
	if stopLoss >= 0 {
		return ErrStopLossIsNotNegative
	}
	err := p.positionRepository.UpdateStopLoss(ctx, userId, positionId, stopLoss)
	if err == repository.ErrPositionNotFound {
		return ErrPositionNotFound
	}
	return err
}

func (p *position) SetTakeProfit(ctx context.Context, userId int64, positionId int64, takeProfit float64) error {
	if takeProfit <= 0 {
		return ErrTakeProfitIsNotPositive
	}
	err := p.positionRepository.UpdateTakeProfit(ctx, userId, positionId, takeProfit)
	if err == repository.ErrPositionNotFound {
		return ErrPositionNotFound
	}
	return err
}

func (p *position) UpdatePortfolio(ctx context.Context, position model.Position) error {
	p.portfolioRepository.UpdatePortfolio(position)
	price, err := p.priceRepository.GetPrice(position.Symbol)
	if err != nil {
		return nil
	}

	return p.RecalculatePortfolio(ctx, position.UserID, position.PositionID, price)
}

func (p *position) RemoveFromPortfolio(position model.Position) {
	p.portfolioRepository.RemoveFromPortfolio(position)
}

func (p *position) RecalculatePortfolio(ctx context.Context, userId, positionId int64, price model.Price) error {
	pnl, err := p.portfolioRepository.RecalculatePortfolio(userId, positionId, price)
	if err != nil {
		return nil
	}

	balance, err := p.userRepository.GetBalance(ctx, userId)
	if err != nil {
		return err
	}

	if balance+pnl <= 0 {
		log.WithFields(log.Fields{
			"user_id": userId,
			"balance": balance,
			"pnl":     pnl,
		}).Info("Closing positions due to insufficient balance")

		positions := p.portfolioRepository.GetAllPositions(userId)
		return p.closePositions(ctx, userId, positions)
	}

	return nil
}

func (p *position) closePositions(ctx context.Context, userId int64, positions []model.Position) error {
	totalProfit := .0
	for _, pos := range positions {
		pr, _ := p.priceRepository.GetPrice(pos.Symbol)
		closePrice := pr.GetPrice(!pos.IsBuyType)
		err := p.positionRepository.ClosePosition(ctx, pos.PositionID, closePrice)
		if err != nil {
			continue
		}

		totalProfit += p.getClosingBalanceChange(pos.IsBuyType, closePrice)
	}
	_, err := p.userRepository.AddToBalance(ctx, userId, totalProfit)
	return err
}

func (p *position) getClosingBalanceChange(isBuyType bool, closePrice float64) float64 {
	if isBuyType {
		return closePrice
	}
	return -closePrice
}
