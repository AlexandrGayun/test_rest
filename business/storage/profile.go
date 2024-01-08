package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"test_task/business/domain"
)

func (s *Storage) GetProfiles(ctx context.Context) ([]domain.Profile, error) {
	res, err := s.queries.GetProfiles(ctx)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("storage: %w", err)
		}
		return nil, nil
	}
	return convertToDomainProfiles(res), nil
}

func (s *Storage) GetProfileByUsername(ctx context.Context, username string) (*domain.Profile, error) {
	res, err := s.queries.GetProfileByUsername(ctx, username)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("storage: %w", err)
		}
		return nil, nil
	}
	return convertUsernameProfileRowToDomainProfile(res), nil
}

func convertToDomainProfiles(input []GetProfilesRow) []domain.Profile {
	res := make([]domain.Profile, len(input))
	for i, p := range input {
		res[i] = convertProfileRowToDomainProfile(p)
	}
	return res
}

func convertProfileRowToDomainProfile(input GetProfilesRow) domain.Profile {
	res := domain.Profile{Id: input.ID, Username: input.Username}
	if input.FirstName.Valid {
		res.FirstName = input.FirstName.String
	}
	if input.LastName.Valid {
		res.LastName = input.LastName.String
	}
	if input.City.Valid {
		res.City = input.City.String
	}
	if input.School.Valid {
		res.School = input.School.String
	}
	return res
}

func convertUsernameProfileRowToDomainProfile(input GetProfileByUsernameRow) *domain.Profile {
	res := domain.Profile{Id: input.ID, Username: input.Username}
	if input.FirstName.Valid {
		res.FirstName = input.FirstName.String
	}
	if input.LastName.Valid {
		res.LastName = input.LastName.String
	}
	if input.City.Valid {
		res.City = input.City.String
	}
	if input.School.Valid {
		res.School = input.School.String
	}
	return &res
}
