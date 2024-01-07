package service

import "go.uber.org/zap"

type Implementor interface {
	CheckAuth(string) error
	GetProfiles()
}

type Service struct {
	logger *zap.Logger
}

func New(logger *zap.Logger) *Service {
	return &Service{logger: logger}
}
