package handler

import (
	"context"
	"github.com/evleria/PriceService/protocol/pb"
)

type BiddingService struct {
	pb.UnimplementedBiddingServiceServer
}

func NewBiddingService() pb.BiddingServiceServer {
	return &BiddingService{}
}

func (s *BiddingService) AddPosition(ctx context.Context, request *pb.AddPositionRequest) (*pb.AddPositionResponse, error) {
	return &pb.AddPositionResponse{
		PositionId: int64(request.InitialPrice),
	}, nil
}
