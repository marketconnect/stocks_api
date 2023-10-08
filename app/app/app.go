package app

import (
	"context"
	"fmt"
	"net"
	"stocks_api/app/internal/config"
	"stocks_api/app/internal/data_provider/auth_data_provider"
	"stocks_api/app/internal/data_provider/card_data_provider"
	"stocks_api/app/internal/data_provider/stock_data_provider"

	auth_service "stocks_api/app/internal/domain/service/auth_service"
	card_service "stocks_api/app/internal/domain/service/card_service"
	"stocks_api/app/internal/domain/service/stock_service"

	"stocks_api/app/pkg/client/postgresql"
	my_jwt "stocks_api/app/pkg/jwt"
	"stocks_api/app/pkg/logger"
	"strconv"
	"time"

	pb "stocks_api/app/gen/proto"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type App struct {
	cfg        *config.Config
	logger     logger.Logger
	grpcServer *grpc.Server
}

func NewApp(config *config.Config, logger logger.Logger) (App, error) {
	logger.Info("Postgres initializing")
	pgConfig := postgresql.NewPgConfig(
		config.PostgreSQL.PostgreUsername, config.PostgreSQL.Password,
		config.PostgreSQL.Host, config.PostgreSQL.Port, config.PostgreSQL.Database,
	)
	pgClient, err := postgresql.NewClient(context.Background(), 5, time.Second*5, pgConfig)
	if err != nil {
		logger.Fatal(err)
	}

	tokenDuration, err := strconv.Atoi(config.Jwt.TokenDuration)
	if err != nil {
		logger.Fatal(err)
	}

	jwtManager := my_jwt.NewJWTManager(config.Jwt.SecretKey, time.Duration((time.Minute * time.Duration(tokenDuration))))

	// Token Manager
	// tokenManager := my_jwt.NewJWTManager(config.Jwt.SecretKey, time.Duration((time.Minute * time.Duration(tokenDuration))))

	// Data Providers
	authDataProvider := auth_data_provider.NewAuthStorage(pgClient)
	stocksDataProvider := stock_data_provider.NewStockStorage(pgClient)
	cardDataProvider := card_data_provider.NewCardStorage(pgClient)

	// Services
	authService := auth_service.NewAuthService(authDataProvider, jwtManager, tokenDuration, logger)
	cardsService := card_service.NewCardService(cardDataProvider, jwtManager, logger)

	stockService := stock_service.NewStockService(stocksDataProvider, logger)
	// interceptor := auth_interceptor.NewAuthInterceptor(subscriptionDataProvider, tokenManager)

	// grpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptor.Unary()))
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, authService)
	pb.RegisterStockServiceServer(grpcServer, stockService)
	pb.RegisterProductCardServiceServer(grpcServer, cardsService)

	return App{
		cfg:        config,
		logger:     logger,
		grpcServer: grpcServer,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	grp, ctx := errgroup.WithContext(ctx)
	grp.Go(func() error {
		return a.startGRPC(ctx)
	})
	return grp.Wait()
}

func (a *App) startGRPC(ctx context.Context) error {
	a.logger.Info("start GRPC")
	address := fmt.Sprintf("%s:%s", a.cfg.GRPC.IP, a.cfg.GRPC.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		a.logger.Fatal("cannot start GRPC server: ", err)
	}
	a.logger.Info("start GRPC server on address %s", address)
	err = a.grpcServer.Serve(listener)
	if err != nil {
		a.logger.Fatal("cannot start GRPC server: ", err)
	}
	return nil
}
