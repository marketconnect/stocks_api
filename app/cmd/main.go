package main

import (
	"context"
	"log"
	"stocks_api/app/app"
	"stocks_api/app/internal/config"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}

	sugar := logger.Sugar()

	sugar.Info("config initializing")
	cfg := config.GetConfig()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	a, err := app.NewApp(ctx, cfg, sugar)
	if err != nil {
		sugar.Fatal(err)
	}

	sugar.Info("Running Application")

	a.Run(ctx)

}
