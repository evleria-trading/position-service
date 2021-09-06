package main

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/evleria-trading/position-service/internal/config"
	"github.com/evleria-trading/position-service/internal/consumer"
	"github.com/evleria-trading/position-service/internal/handler"
	"github.com/evleria-trading/position-service/internal/pnl"
	"github.com/evleria-trading/position-service/internal/repository"
	"github.com/evleria-trading/position-service/internal/service"
	"github.com/evleria-trading/position-service/protocol/pb"
	pricePb "github.com/evleria/price-service/protocol/pb"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	cfg := new(config.Сonfig)
	err := env.Parse(cfg)
	if err != nil {
		log.Fatal(err)
	}

	setupLogger(cfg.Environment)

	db := getPostgres(cfg)
	defer db.Close()

	priceClient := getPriceGrpcClient(cfg)

	positionRepository := repository.NewPositionRepository(db)
	priceRepository := repository.NewPriceRepository()
	positionService := service.NewPositionService(positionRepository, priceRepository)
	priceConsumer := consumer.NewPriceConsumer(priceClient, priceRepository)
	pricesChan, err := priceConsumer.Consume(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	openedChan, closedChan, updatedChan, err := positionRepository.ListenNotifications(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	pnlMonitor := pnl.NewMonitor(positionService)
	go pnlMonitor.CalculatePnlForOpenPositions(pricesChan, openedChan, closedChan, updatedChan)

	startGrpcServer(handler.NewPositionService(positionService), ":6000")
}

func startGrpcServer(positionService pb.PositionServiceServer, port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterPositionServiceServer(s, positionService)
	reflection.Register(s)

	log.Info("gRPC server started on ", port)
	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}

func setupLogger(environment string) {
	switch environment {
	case "prod":
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.InfoLevel)
	default:
		log.SetLevel(log.DebugLevel)
	}
}

func getPostgres(cfg *config.Сonfig) *pgxpool.Pool {
	dbURL := getPostgresConnectionString(cfg)

	db, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func getPriceGrpcClient(cfg *config.Сonfig) pricePb.PriceServiceClient {
	url := fmt.Sprintf("%s:%d", cfg.PriceServiceHost, cfg.PriceServicePort)
	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	return pricePb.NewPriceServiceClient(conn)
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
