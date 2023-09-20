package app

import (
	"context"
	"fmt"
	"net"
	"stocks_api/app/internal/config"
	"stocks_api/app/internal/data_provider/auth_data_provider"
	"stocks_api/app/internal/data_provider/subscription_data_provider"
	"stocks_api/app/internal/domain/service/auth_interceptor"
	auth_service "stocks_api/app/internal/domain/service/auth_service"
	"stocks_api/app/internal/domain/service/subscription_service"
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

func NewApp(ctx context.Context, config *config.Config, logger logger.Logger) (App, error) {
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
	tokenManager := my_jwt.NewJWTManager(config.Jwt.SecretKey, time.Duration((time.Minute * time.Duration(tokenDuration))))

	// Data Providers
	authDataProvider := auth_data_provider.NewAuthStorage(pgClient)
	subscriptionDataProvider := subscription_data_provider.NewSubscriptionStorage(pgClient)

	// Services
	authService := auth_service.NewAuthService(authDataProvider, subscriptionDataProvider, jwtManager, logger)
	subscriptionService := subscription_service.NewSubscriptionService(subscriptionDataProvider, logger)
	interceptor := auth_interceptor.NewAuthInterceptor(subscriptionDataProvider, tokenManager)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptor.Unary()))
	pb.RegisterAuthServiceServer(grpcServer, authService)
	pb.RegisterUserSubscriptionsServiceServer(grpcServer, subscriptionService)

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
