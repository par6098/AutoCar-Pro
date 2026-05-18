package internal

import (
	"context"
	"log"
	"time"

	clickhouse "github.com/ClickHouse/clickhouse-go/v2"
)

func ConnectClickHouse(cfg Config) clickhouse.Conn {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{cfg.ClickHouseAddr},
		Auth: clickhouse.Auth{
			Database: cfg.ClickHouseDB,
			Username: cfg.ClickHouseUser,
			Password: cfg.ClickHousePassword,
		},
		DialTimeout:     10 * time.Second,
		ConnMaxLifetime: time.Hour,
		MaxOpenConns:    10,
		MaxIdleConns:    5,
	})

	if err != nil {
		log.Fatal("ClickHouse connection failed: ", err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		log.Fatal("ClickHouse ping failed: ", err)
	}

	return conn
}
