package handler

import (
	"context"
	"github.com/evleria-trading/position-service/internal/service"
	"github.com/evleria-trading/position-service/protocol/pb"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/guregu/null.v4"
)

type PositionService struct {
	pb.UnimplementedPositionServiceServer
	service service.Position
}

func NewPositionService(positionService service.Position) pb.PositionServiceServer {
	return &PositionService{
		service: positionService,
	}
}

func (p *PositionService) OpenPosition(ctx context.Context, request *pb.OpenPositionRequest) (*pb.OpenPositionResponse, error) {
	id, err := p.service.AddPosition(ctx, request.UserId, request.Symbol, request.IsBuyType, request.PriceId)
	if err != nil {
		return nil, status.Error(getStatusCode(err), err.Error())
	}

	return &pb.OpenPositionResponse{
		PositionId: id,
	}, nil
}

func (p *PositionService) ClosePosition(ctx context.Context, request *pb.ClosePositionRequest) (*pb.ClosePositionResponse, error) {
	profit, err := p.service.ClosePosition(ctx, request.UserId, request.PositionId, request.PriceId)
	if err != nil {
		return nil, status.Error(getStatusCode(err), err.Error())
	}

	return &pb.ClosePositionResponse{
		Profit: profit,
	}, nil
}

func (p *PositionService) GetOpenPosition(ctx context.Context, request *pb.GetOpenPositionRequest) (*pb.GetOpenPositionResponse, error) {
	pos, err := p.service.GetOpenPosition(ctx, request.PositionId)
	if err != nil {
		return nil, status.Error(getStatusCode(err), err.Error())
	}

	return &pb.GetOpenPositionResponse{
		AddPrice:   pos.AddPrice,
		Symbol:     pos.Symbol,
		IsBuyType:  pos.IsBuyType,
		StopLoss:   toProtoDoubleValue(pos.StopLoss),
		TakeProfit: toProtoDoubleValue(pos.TakeProfit),
		UserId:     pos.UserID,
	}, nil
}

func (p *PositionService) SetStopLoss(ctx context.Context, request *pb.SetStopLossRequest) (*empty.Empty, error) {
	err := p.service.SetStopLoss(ctx, request.UserId, request.PositionId, request.StopLoss)
	if err != nil {
		return nil, status.Error(getStatusCode(err), err.Error())
	}
	return &empty.Empty{}, nil
}

func (p *PositionService) SetTakeProfit(ctx context.Context, request *pb.SetTakeProfitRequest) (*empty.Empty, error) {
	err := p.service.SetTakeProfit(ctx, request.UserId, request.PositionId, request.TakeProfit)
	if err != nil {
		return nil, status.Error(getStatusCode(err), err.Error())
	}
	return &empty.Empty{}, nil
}

func getStatusCode(err error) codes.Code {
	switch err {
	case service.ErrPriceNotFound, service.ErrPositionNotFound:
		return codes.NotFound
	case service.ErrPositionClosed:
		return codes.FailedPrecondition
	case service.ErrPriceChanged, service.ErrStopLossIsNotNegative:
		return codes.InvalidArgument
	default:
		return codes.Internal
	}
}

func toProtoDoubleValue(f null.Float) *wrappers.DoubleValue {
	if f.Valid == false {
		return nil
	}
	return &wrappers.DoubleValue{
		Value: f.Float64,
	}
}
