package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func (s *Storage) GetApiKeyID(ctx context.Context, apikey string) (*int64, error) {
	res, err := s.queries.GetApiKeyID(ctx, apikey)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("storage: %w", err)
		} else {
			return nil, nil
		}
	}
	return &res, nil
}
