package postgres

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	clt "github.com/lefties-startup/companion-api/internal/client"
)

type PostgresConnection interface { //nolint:revive
	ExecContext(ctx context.Context, telemetryName string, query string, args ...any) (sql.Result, error)
	GetContext(ctx context.Context, telemetryName string, dest any, query string, args ...any) error
	NamedExecContext(ctx context.Context, telemetryName string, query string, arg any) (sql.Result, error)
	PrepareNamedContext(ctx context.Context, telemetryName string, query string) (*sqlx.NamedStmt, error)
	SelectContext(ctx context.Context, telemetryName string, dest any, query string, args ...any) error
}

type Client struct {
	connection PostgresConnection
}

func New(connection PostgresConnection) (*Client, error) {
	if connection == nil {
		return nil, clt.ErrConnectionIsNil
	}

	return &Client{connection: connection}, nil
}

func (c *Client) Exec(ctx context.Context, telemetryName string, query string, args ...any) error {
	if ctx == nil {
		return errors.Wrap(clt.ErrArgumentIsNil, "ctx")
	}

	_, err := c.connection.ExecContext(ctx, telemetryName, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) NamedExec(ctx context.Context, telemetryName string, query string, arg any) error {
	if ctx == nil {
		return errors.Wrap(clt.ErrArgumentIsNil, "ctx")
	}

	_, err := c.connection.NamedExecContext(ctx, telemetryName, query, arg)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) NamedQuery(ctx context.Context, telemetryName string, result any, query string, arg any) error {
	if ctx == nil {
		return errors.Wrap(clt.ErrArgumentIsNil, "ctx")
	}

	stmt, err := c.connection.PrepareNamedContext(ctx, telemetryName, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.GetContext(ctx, result, arg)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Row(ctx context.Context, telemetryName string, result any, query string, args ...any) error {
	if ctx == nil {
		return errors.Wrap(clt.ErrArgumentIsNil, "ctx")
	}
	if args == nil {
		return errors.Wrap(clt.ErrArgumentIsNil, "args")
	}

	err := c.connection.GetContext(ctx, telemetryName, result, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) Rows(ctx context.Context, telemetryName string, result any, query string, args ...any) error {
	if ctx == nil {
		return errors.Wrap(clt.ErrArgumentIsNil, "ctx")
	}
	if args == nil {
		return errors.Wrap(clt.ErrArgumentIsNil, "args")
	}

	err := c.connection.SelectContext(ctx, telemetryName, result, query, args...)
	if err != nil {
		return err
	}

	return nil
}
