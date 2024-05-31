package database

import (
	"context"
	"fmt"
	"log"

	"github.com/elhaqeeem/postgres-Golang-htmx/settings"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

// NewPostgresDB creates a new Postgres database connection
func NewPostgresDB(
	ctx context.Context, s *settings.Settings,
) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		s.DB.Host,
		s.DB.Port,
		s.DB.User,
		s.DB.Password,
		s.DB.Name,
	)

	db, err := sqlx.ConnectContext(ctx, "postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("🔥 failed to connect to the database: %s", err)
	}

	// sqlx.BindDriver("postgres", sqlx.DOLLAR)

	log.Println("🚀 Connected Successfully to the Database")

	return db, nil
}
