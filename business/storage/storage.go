package storage

import "go.uber.org/zap"

type Storage struct {
	queries *Queries
	logger  *zap.Logger // could define a different logger for a db layer
}

func NewStorage(q *Queries, l *zap.Logger) *Storage {
	return &Storage{queries: q, logger: l}
}
