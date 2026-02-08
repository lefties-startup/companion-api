package user

import (
	"context"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	modelDB "github.com/lefties-startup/companion-api/internal/model/db"
	"github.com/lefties-startup/companion-api/internal/service"
)

// GetInfo Get: /api/v1/me/{user_id}
func (s *ServiceUser) GetInfo(ctx context.Context, userID int64) (*modelDB.User, error) {
	log := s.logger.
		With(zap.String("method", "ServiceUser.GetInfo")).
		With(zap.Any("userID", userID))

	user, err := s.repository.GetUser(ctx, userID)
	if errors.Is(err, service.NotFoundErr) {
		log.
			With(zap.String("returned_error", err.Error())).
			Error("User not found")

		return nil, errors.Wrap(err, "not found user info")
	}
	if err != nil {
		log.
			With(zap.String("internal_error", err.Error())).
			Error("Error get user info")

		return nil, errors.Wrap(err, "error get user info")
	}
	return user, nil
}

// PatchInfo Get: /api/v1/me/{user_id}
func (s *ServiceUser) PatchInfo(ctx context.Context, user modelDB.User) error {
	log := s.logger.
		With(zap.String("method", "ServiceUser.GetInfo")).
		With(zap.Any("user", user))

	err := s.repository.UpsertUser(ctx, &user)
	if err != nil {
		log.
			With(zap.String("internal_error", err.Error())).
			Error("Error upsert user info")

		return errors.Wrap(err, "error upsert user")
	}
	return nil
}

// CreateUser Post: /api/v1/auth/register
func (s *ServiceUser) CreateUser(ctx context.Context, user modelDB.User) error {
	log := s.logger.
		With(zap.String("method", "ServiceUser.GetInfo")).
		With(zap.Any("user", user))

	err := s.repository.UpsertUser(ctx, &user)
	if err != nil {
		log.
			With(zap.String("internal_error", err.Error())).
			Error("Error create user")

		return errors.Wrap(err, "error create user")
	}
	return nil
}
