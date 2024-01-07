package handler

import (
	"go.uber.org/zap"
	"test_task/business/service"
)

type Handler struct {
	logger  *zap.Logger
	service service.Implementor
}

func New(logger *zap.Logger, service service.Implementor) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
	}
}
