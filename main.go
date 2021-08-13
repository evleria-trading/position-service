package main

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/evleria/PriceService/internal/config"
	"github.com/evleria/PriceService/internal/generator"
	grpcService "github.com/evleria/PriceService/internal/handler"
	"github.com/evleria/PriceService/internal/producer"
	"github.com/evleria/PriceService/internal/repository"
	"github.com/evleria/PriceService/internal/service"
	"github.com/evleria/PriceService/protocol/pb"
	"github.com/go-redis/redis/v8"
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

	redisClient := getRedis(cfg)
	defer redisClient.Close()

	positionRepository := repository.NewPositionService(db)
	priceProducer := producer.NewProducerPrice(redisClient)
	priceGenerator := generator.NewPricesGenerator(priceProducer)
	positionService := service.NewPositionService(positionRepository, priceProducer)

	if cfg.GeneratePrices {
		go func() {
			err := priceGenerator.GeneratePrices(context.Background())
			check(err)
		}()
	}

	biddingService := grpcService.NewBiddingService(positionService)
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

func getRedis(cfg *config.Сonfig) *redis.Client {
	opts := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPass,
	}

	redisClient := redis.NewClient(opts)
	_, err := redisClient.Ping(context.Background()).Result()
	check(err)

	return redisClient
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
