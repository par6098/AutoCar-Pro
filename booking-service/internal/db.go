package internal

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(databaseURL string) *pgxpool.Pool {
	db, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	if err := db.Ping(context.Background()); err != nil {
		log.Fatal("DB ping failed:", err)
	}

	return db
}
