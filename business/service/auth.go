package service

import (
	"context"
	"fmt"
)

func (s *Service) CheckAuth(ctx context.Context, apikey string) (*int64, error) {
	res, err := s.storage.GetApiKeyID(ctx, apikey)
	if err != nil {
		return nil, fmt.Errorf("service: %w", err)
	}
	return res, nil
}
