package internal

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ClickHouseAddr     string
	ClickHouseDB       string
	ClickHouseUser     string
	ClickHousePassword string

	RedisAddr         string
	RedisPassword     string
	RedisDB           int
	BookingEventQueue string

	JWTSecret string
}

func LoadConfig() Config {
	_ = godotenv.Load()

	redisDB, _ := strconv.Atoi(getEnv("REDIS_DB", "0"))

	return Config{
		ClickHouseAddr:     getEnv("CLICKHOUSE_ADDR", "localhost:9000"),
		ClickHouseDB:       getEnv("CLICKHOUSE_DB", "autocare_analytics"),
		ClickHouseUser:     getEnv("CLICKHOUSE_USER", "default"),
		ClickHousePassword: getEnv("CLICKHOUSE_PASSWORD", ""),

		RedisAddr:         getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword:     getEnv("REDIS_PASSWORD", ""),
		RedisDB:           redisDB,
		BookingEventQueue: getEnv("BOOKING_EVENT_QUEUE", "booking-events"),

		JWTSecret: getEnv("JWT_SECRET", "autocare_secret_key"),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
