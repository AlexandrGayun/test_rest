package service

import (
	"context"
	"fmt"
	"test_task/business/domain"
)

func (s *Service) GetProfiles(ctx context.Context) ([]domain.Profile, error) {
	res, err := s.storage.GetProfiles(ctx)
	if err != nil {
		return nil, fmt.Errorf("service: %w", err)
	}
	return res, nil
}
func (s *Service) GetProfileByUsername(ctx context.Context, username string) (*domain.Profile, error) {
	res, err := s.storage.GetProfileByUsername(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("service: %w", err)
	}
	return res, nil
}
