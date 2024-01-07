package logger

import (
	"go.uber.org/zap"
	"log"
)

func New() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)
	return logger
}
