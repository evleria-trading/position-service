package main

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/evleria/PriceService/internal/config"
	"github.com/evleria/PriceService/internal/consumer"
	"github.com/evleria/PriceService/internal/generator"
	grpcService "github.com/evleria/PriceService/internal/handler"
	"github.com/evleria/PriceService/internal/producer"
	"github.com/evleria/PriceService/internal/repository"
	"github.com/evleria/PriceService/internal/service"
	"github.com/evleria/PriceService/protocol/pb"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

func main() {
	cfg := new(config.Сonfig)
	check(env.Parse(cfg))

	setupLogger(cfg.Environment)

	db := getPostgres(cfg)
	defer db.Close(context.Background())

	redisClient := getRedis(cfg)
	defer redisClient.Close()

	positionRepository := repository.NewPositionService(db)
	priceRepository := repository.NewPriceRepository()
	priceProducer := producer.NewProducerPrice(redisClient)
	priceGenerator := generator.NewPricesGenerator(priceProducer, cfg.GenerationRate)
	positionService := service.NewPositionService(positionRepository, priceRepository)

	if cfg.GeneratePrices {
		go func() {
			err := priceGenerator.GeneratePrices(context.Background())
			check(err)
		}()
	}

	priceConsumer := consumer.NewPriceConsumer(redisClient, priceRepository, cfg.ConsumerWarmup)
	go func() {
		err := priceConsumer.Consume(context.Background())
		check(err)
	}()

	biddingService := grpcService.NewBiddingService(positionService)
	startGrpcServer(biddingService, ":6000")
}

func startGrpcServer(biddingService pb.PositionServiceServer, port string) {
	listener, err := net.Listen("tcp", port)
	check(err)

	s := grpc.NewServer()
	pb.RegisterPositionServiceServer(s, biddingService)
	reflection.Register(s)

	log.Info("gRPC server started on ", port)
	check(s.Serve(listener))
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
