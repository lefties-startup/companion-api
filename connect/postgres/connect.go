package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v5/stdlib"
	"net"

	"github.com/jackc/pgx/v5"
	"github.com/jmoiron/sqlx"
)

type ctxTelemetryKey int

const ctxTelemetryName ctxTelemetryKey = iota

//type TransactionFunc func(ctx context.Context, tx *Tx) error

// Config is a connection configuration.
type Config struct {
	Host          string `mapstructure:"host"`
	Port          string `mapstructure:"port"`
	DB            string `mapstructure:"db"`
	User          string `mapstructure:"user"`
	Password      string `mapstructure:"password"`
	MaxConnection int    `mapstructure:"max_connection"`
}

func (c Config) ConnURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s/%s", c.User, c.Password, net.JoinHostPort(c.Host, c.Port), c.DB)
}

type Connection struct {
	db *sqlx.DB
	//tx *sqlx.Tx
}

func OpenConnection(config Config) (*Connection, error) {
	db, err := open(config)
	if err != nil {
		return nil, err
	}

	return &Connection{db: db}, nil
}

func (c *Connection) Ping(ctx context.Context) error {
	return c.db.PingContext(ctx)
}

func (c *Connection) Close() error {
	return c.db.Close()
}

func open(config Config) (*sqlx.DB, error) {
	pgxConfig, err := pgx.ParseConfig(config.ConnURL())
	if err != nil {
		return nil, err
	}

	conn := sqlx.NewDb(stdlib.OpenDB(*pgxConfig), "pgx")
	conn.SetMaxOpenConns(config.MaxConnection)
	if err = conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}

func (c *Connection) ExecContext(ctx context.Context, telemetryName string, query string, args ...any) (sql.Result, error) {
	return c.db.ExecContext(context.WithValue(ctx, ctxTelemetryName, telemetryName), query, args...)
}

func (c *Connection) GetContext(ctx context.Context, telemetryName string, dest any, query string, args ...any) error {
	return c.db.GetContext(context.WithValue(ctx, ctxTelemetryName, telemetryName), dest, query, args...)
}

func (c *Connection) NamedExecContext(ctx context.Context, telemetryName string, query string, arg any) (sql.Result, error) {
	return c.db.NamedExecContext(context.WithValue(ctx, ctxTelemetryName, telemetryName), query, arg)
}

func (c *Connection) PrepareNamedContext(ctx context.Context, telemetryName string, query string) (*sqlx.NamedStmt, error) {
	return c.db.PrepareNamedContext(context.WithValue(ctx, ctxTelemetryName, telemetryName), query)
}

func (c *Connection) SelectContext(ctx context.Context, telemetryName string, dest any, query string, args ...any) error {
	return c.db.SelectContext(context.WithValue(ctx, ctxTelemetryName, telemetryName), dest, query, args...)
}
