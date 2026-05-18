package internal

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(databaseURL string) *pgxpool.Pool {

	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Fatal("failed to parse database config: ", err)
	}

	config.MaxConns = 20
	config.MinConns = 5
	config.MaxConnLifetime = time.Hour
	config.MaxConnIdleTime = 30 * time.Minute
	config.HealthCheckPeriod = time.Minute

	db, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal("failed to create connection pool: ", err)
	}

	if err := db.Ping(context.Background()); err != nil {
		log.Fatal("failed to ping database: ", err)
	}

	log.Println("PostgreSQL connected successfully")

	return db
}
