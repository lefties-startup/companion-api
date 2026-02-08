package user

import (
	"context"
	modelDB "github.com/lefties-startup/companion-api/internal/model/db"
	"github.com/lefties-startup/companion-api/internal/service"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Repository interface {
	GetUser(ctx context.Context, userID int64) (*modelDB.User, error)
	UpsertUser(ctx context.Context, user *modelDB.User) error
}
type ServiceUser struct {
	logger     *zap.Logger
	repository Repository
}

func NewServiceUser(logger *zap.Logger, repository Repository) (*ServiceUser, error) {
	if logger == nil {
		return nil, errors.Wrap(service.NilLoggerErr, "NewServiceUser")
	}
	if repository == nil {
		return nil, errors.Wrap(service.NilRepositoryErr, "NewServiceUser")
	}
	return &ServiceUser{
		logger:     logger,
		repository: repository,
	}, nil
}
