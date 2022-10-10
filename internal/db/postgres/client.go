package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type Client struct {
	db *pgx.Conn
}

func NewClient(connectionString string) (*Client, error) {
	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}
	return &Client{db: conn}, nil
}
