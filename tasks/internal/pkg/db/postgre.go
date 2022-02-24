package db

import (
	"context"
	"database/sql"
	"fmt"
)

func PostgreSQLConnect(ctx context.Context, dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to load driver: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to load driver: %w", err)
	}
	return db, nil
}

func Close(ctx context.Context, db *sql.DB) error {
	return db.Close()
}
