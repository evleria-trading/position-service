package main

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/evleria/position-service/internal/config"
	"github.com/evleria/position-service/internal/consumer"
	grpcService "github.com/evleria/position-service/internal/handler"
	"github.com/evleria/position-service/internal/repository"
	"github.com/evleria/position-service/internal/service"
	"github.com/evleria/position-service/protocol/pb"
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
	positionService := service.NewPositionService(positionRepository, priceRepository)

	priceConsumer := consumer.NewPriceConsumer(redisClient, priceRepository, cfg.ConsumerWarmup)
	go func() {
		err := priceConsumer.Consume(context.Background())
		check(err)
	}()

	startGrpcServer(grpcService.NewPositionService(positionService), ":6000")
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
