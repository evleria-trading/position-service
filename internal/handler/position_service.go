package handler

import (
	"context"
	"github.com/evleria/PriceService/internal/service"
	"github.com/evleria/PriceService/protocol/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PositionService struct {
	pb.UnimplementedPositionServiceServer
	service service.Position
}

func NewBiddingService(positionService service.Position) pb.PositionServiceServer {
	return &PositionService{
		service: positionService,
	}
}

func (p *PositionService) OpenPosition(ctx context.Context, request *pb.OpenPositionRequest) (*pb.OpenPositionResponse, error) {
	id, err := p.service.AddPosition(ctx, request.Symbol, request.IsBuyType)
	if err != nil {
		if err == service.ErrPriceNotFound {
			return nil, status.Error(codes.NotFound, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.OpenPositionResponse{
		PositionId: int64(id),
	}, nil
}
