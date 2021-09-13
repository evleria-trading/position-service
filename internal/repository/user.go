package repository

import (
	"context"
	userPb "github.com/evleria-trading/user-service/protocol/pb"
	log "github.com/sirupsen/logrus"
)

type User interface {
	GetBalance(ctx context.Context, userId int64) (float64, error)
	SetBalance(ctx context.Context, userId int64, balance float64) error
	AddToBalance(ctx context.Context, id int64, profit float64) (float64, error)
}

type user struct {
	client userPb.UserServiceClient
}

func NewUserRepository(client userPb.UserServiceClient) User {
	return &user{
		client: client,
	}
}

func (u *user) GetBalance(ctx context.Context, userId int64) (float64, error) {
	request := &userPb.GetBalanceRequest{
		UserId: userId,
	}
	response, err := u.client.GetBalance(ctx, request)
	if err != nil {
		return 0, err
	}
	return response.Balance, nil
}

func (u *user) SetBalance(ctx context.Context, userId int64, balance float64) error {
	request := &userPb.SetBalanceRequest{
		UserId:  userId,
		Balance: balance,
	}

	_, err := u.client.SetBalance(ctx, request)
	return err
}

func (u *user) AddToBalance(ctx context.Context, id int64, profit float64) (float64, error) {
	request := &userPb.AddToBalanceRequest{
		UserId: id,
		Diff:   profit,
	}

	response, err := u.client.AddToBalance(ctx, request)
	if err != nil {
		return 0, err
	}

	log.WithFields(log.Fields{
		"user_id":     id,
		"diff":        profit,
		"new_balance": response.Balance,
	}).Info("Successfully updated user balance")
	return response.Balance, nil
}
