package repository

import (
	"context"
	userPb "github.com/evleria-trading/user-service/protocol/pb"
)

type User interface {
	GetBalance(ctx context.Context, userId int64) (float64, error)
	SetBalance(ctx context.Context, userId int64, balance float64) error
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
