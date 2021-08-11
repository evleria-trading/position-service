package main

import (
  "fmt"
  grpcService "gh/evleria/PriceService/internal/grpc"
  "gh/evleria/PriceService/protocol/pb"
  "google.golang.org/grpc"
  "google.golang.org/grpc/reflection"
  "log"
  "net"
)

func main() {
    fmt.Println("Hello, World")
    biddingService := grpcService.NewBiddingService()
    startGrpcServer(biddingService, ":6000")
}

func startGrpcServer(biddingService pb.BiddingServiceServer, port string) {
  listener, err := net.Listen("tcp", port)
  check(err)

  s := grpc.NewServer()
  pb.RegisterBiddingServiceServer(s,biddingService)
  reflection.Register(s)

  check(s.Serve(listener))
}

func check(err error) {
  if err != nil {
    log.Fatal(err)
  }
}