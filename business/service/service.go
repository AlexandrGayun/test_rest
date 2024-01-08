package service

import (
	"context"
	"go.uber.org/zap"
	"test_task/business/domain"
	"test_task/business/storage"
)

type Implementor interface {
	CheckAuth(context.Context, string) (*int64, error)
	GetProfiles(context.Context) ([]domain.Profile, error)
	GetProfileByUsername(context.Context, string) (*domain.Profile, error)
}

type Service struct {
	logger  *zap.Logger
	storage *storage.Storage
}

func New(logger *zap.Logger, storage *storage.Storage) *Service {
	return &Service{logger: logger, storage: storage}
}
