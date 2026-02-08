package main

import (
	"context"
	"github.com/lefties-startup/companion-api/connect/http"
	"github.com/lefties-startup/companion-api/connect/postgres"
	"github.com/lefties-startup/companion-api/internal/companion"
	deliveryHTTP "github.com/lefties-startup/companion-api/internal/delivery/http"
	userServ "github.com/lefties-startup/companion-api/internal/service/user"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	pgClient "github.com/lefties-startup/companion-api/internal/client/postgres"
	pgRepository "github.com/lefties-startup/companion-api/internal/repository/postgres"
)

func main() {
	loggingLevel := zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	loggingConfig := zap.NewProductionConfig()
	loggingConfig.Level = loggingLevel
	loggingConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	logger := zap.Must(loggingConfig.Build())

	defer logger.Sync() //nolint:errcheck
	defer catchPanic(logger)

	config, err := companion.Load()
	if err != nil {
		logger.Error("error to parse configuration", zap.Error(err))
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Создаём сервер
	opsSrv := http.NewOpsServer(":8080", logger)

	pgConn, err := postgres.OpenConnection(config.Postgres.Manager)
	if err != nil {
		logger.Error("error to open postgres master connection", zap.Error(err))
		return
	}
	postgresClient, err := pgClient.New(pgConn)
	if err != nil {
		logger.Error("error to create postgres client", zap.Error(err))
		return
	}

	userRepository, err := pgRepository.NewRepository(postgresClient)
	if err != nil {
		logger.Error("error to create category postgres repository", zap.Error(err))
		return
	}

	userService, err := userServ.NewServiceUser(logger, userRepository)
	if err != nil {
		if err != nil {
			logger.Error("error to create category cache repository", zap.Error(err))
			return
		}
	}
	// Создаём хендлеры
	userHandler := deliveryHTTP.NewUserHandler(*userService, logger)

	// Регистрируем роуты
	opsSrv.Register("GET /api/v1/users/info", userHandler.GetInfo)
	if err := opsSrv.Run(ctx); err != nil {
		logger.Fatal("server failed", zap.Error(err))
	}
}

func catchPanic(logger *zap.Logger) {
	if rec := recover(); rec != nil {
		// Just care about structured logs.
		logger.Panic("caught panic", zap.Any("panic", rec))
	}
}
