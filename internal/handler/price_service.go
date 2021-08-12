package handler

import (
	"context"
	"github.com/evleria/PriceService/internal/service"
	"github.com/evleria/PriceService/protocol/pb"
)

type BiddingService struct {
	pb.UnimplementedBiddingServiceServer
	positionService service.Position
}

func NewBiddingService(positionService service.Position) pb.BiddingServiceServer {
	return &BiddingService{
		positionService: positionService,
	}
}

func (s *BiddingService) AddPosition(ctx context.Context, request *pb.AddPositionRequest) (*pb.AddPositionResponse, error) {
	return &pb.AddPositionResponse{
		PositionId: int64(request.InitialPrice),
	}, nil
}
