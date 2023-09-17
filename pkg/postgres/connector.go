package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// NewConnector provides new PostgreSQL connector
func NewConnector(ctx context.Context, config Config) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, config.ConnectionString)
	if err != nil {
		return nil, fmt.Errorf("postgres connector: %w", err)
	}
	return conn, nil
}
