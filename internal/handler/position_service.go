package handler

import (
	"context"
	"github.com/evleria/position-service/internal/service"
	"github.com/evleria/position-service/protocol/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	id, err := p.service.AddPosition(ctx, request.Symbol, request.IsBuyType, request.PriceId)
	if err != nil {
		switch err {
		case service.ErrPriceNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		case service.ErrPriceChanged:
			return nil, status.Error(codes.InvalidArgument, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &pb.OpenPositionResponse{
		PositionId: id,
	}, nil
}

func (p *PositionService) ClosePosition(ctx context.Context, request *pb.ClosePositionRequest) (*pb.ClosePositionResponse, error) {
	profit, err := p.service.ClosePosition(ctx, request.PositionId, request.PriceId)
	if err != nil {
		switch err {
		case service.ErrPriceNotFound, service.ErrPositionNotFound:
			return nil, status.Error(codes.NotFound, err.Error())
		case service.ErrPositionClosed:
			return nil, status.Error(codes.FailedPrecondition, err.Error())
		case service.ErrPriceChanged:
			return nil, status.Error(codes.InvalidArgument, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &pb.ClosePositionResponse{
		Profit: profit,
	}, nil
}
