package main

import (
	"context"
	"github.com/lefties-startup/companion-api/connect/http"
	deliveryHTTP "github.com/lefties-startup/companion-api/internal/delivery/http"
	userServ "github.com/lefties-startup/companion-api/internal/service/user"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	loggingLevel := zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	loggingConfig := zap.NewProductionConfig()
	loggingConfig.Level = loggingLevel
	loggingConfig.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	logger := zap.Must(loggingConfig.Build())

	defer logger.Sync() //nolint:errcheck
	defer catchPanic(logger)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Создаём сервер
	opsSrv := http.NewOpsServer(":8080", logger)

	userService, err := userServ.NewServiceUser(logger)
	// Создаём хендлеры
	userHandler := deliveryHTTP.NewUserHandler(*userService, logger)

	// Регистрируем роуты
	opsSrv.Register("GET /api/v1/users/info", userHandler.GetInfo)
}

func catchPanic(logger *zap.Logger) {
	if rec := recover(); rec != nil {
		// Just care about structured logs.
		logger.Panic("caught panic", zap.Any("panic", rec))
	}
}
