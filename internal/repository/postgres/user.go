package postgres

import (
	"context"
	"database/sql"
	modelDB "github.com/lefties-startup/companion-api/internal/model/db"
	"github.com/lefties-startup/companion-api/internal/repository"
	"github.com/pkg/errors"
)

type UserClient interface {
	Exec(ctx context.Context, telemetryName string, query string, args ...any) error
	Row(ctx context.Context, telemetryName string, result any, query string, args ...any) error
	Rows(ctx context.Context, telemetryName string, result any, query string, args ...any) error
	NamedExec(ctx context.Context, telemetryName string, query string, arg any) error
	NamedQuery(ctx context.Context, telemetryName string, result any, query string, arg any) error
}

type UserRepository struct {
	client UserClient
}

func NewRepository(client UserClient) (*UserRepository, error) {
	if client == nil {
		return nil, errors.Wrap(repository.ErrClientIsNil, "UserClient")
	}

	return &UserRepository{
		client: client,
	}, nil
}

func (u *UserRepository) GetUser(ctx context.Context, userID int64) (*modelDB.User, error) {
	var user modelDB.User
	err := u.client.Row(
		ctx,
		"UserRepository.GetUser",
		&user,
		`SELECT id, tg_user_id, name, role, created_at, updated_at
		FROM users
		WHERE tg_user_id = $1
		LIMIT 1;`,
		userID,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil //nolint:nilnil
	}
	if err != nil {
		return nil, errors.Wrap(err, "Row")
	}

	return &user, nil
}

func (u *UserRepository) UpsertUser(ctx context.Context, user *modelDB.User) error {
	return nil
}
