package grpc

import (
	"context"
	"github.com/evleria/PriceService/protocol/pb"
)

type BiddingService struct {
	pb. UnimplementedBiddingServiceServer
}

func NewBiddingService() pb.BiddingServiceServer {
return &BiddingService{}
}

func (s *BiddingService) Hello(ctx context.Context,request *pb.HelloRequest) (*pb.HelloResponse, error) {
	response := &pb.HelloResponse{
		Name: request.Name ,
	}
	return response,nil
}
