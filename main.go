package main

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/evleria/PriceService/internal/config"
	grpcService "github.com/evleria/PriceService/internal/grpc"
	"github.com/evleria/PriceService/protocol/pb"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	cfg := new(config.Сonfig)
	check(env.Parse(cfg))

	db := getPostgres(cfg)
	defer db.Close(context.Background())

	//fmt.Println("Hello, World")
	biddingService := grpcService.NewBiddingService()
	startGrpcServer(biddingService, ":6000")
}

func startGrpcServer(biddingService pb.BiddingServiceServer, port string) {
	listener, err := net.Listen("tcp", port)
	check(err)

	s := grpc.NewServer()
	pb.RegisterBiddingServiceServer(s, biddingService)
	reflection.Register(s)

	check(s.Serve(listener))
}

func getPostgres(cfg *config.Сonfig) *pgx.Conn {
	dbURL := getPostgresConnectionString(cfg)
	db, err := pgx.Connect(context.Background(), dbURL)
	check(err)

	return db
}

func getPostgresConnectionString(cfg *config.Сonfig) string {
	conn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.PostgresUser,
		cfg.PostgresPass,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDb,
	)
	if cfg.PostgresSSLDisable {
		conn += "?sslmode=disable"
	}
	return conn
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
