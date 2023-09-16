package app

import (
	"context"
	"fmt"
	"net"
	"stocks_api/app/internal/config"
	"stocks_api/app/pkg/logger"

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
	// pgConfig := postgresql.NewPgConfig(
	// 	config.PostgreSQL.PostgreUsername, config.PostgreSQL.Password,
	// 	config.PostgreSQL.Host, config.PostgreSQL.Port, config.PostgreSQL.Database,
	// )
	// pgClient, err := postgresql.NewClient(context.Background(), 5, time.Second*5, pgConfig)
	// if err != nil {
	// 	logger.Fatal(err)
	// }
	grpcServer := grpc.NewServer()
	return App{
		cfg:        config,
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
