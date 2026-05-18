package internal

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL        string
	JWTSecret          string
	AccessTokenMinutes int
	RefreshTokenDays   int
}

func LoadConfig() Config {

	// Load .env for local development
	_ = godotenv.Load()

	accessMinutes := getEnvAsInt("ACCESS_TOKEN_MINUTES", 15)
	refreshDays := getEnvAsInt("REFRESH_TOKEN_DAYS", 30)

	cfg := Config{
		DatabaseURL:        getEnv("DATABASE_URL", ""),
		JWTSecret:          getEnv("JWT_SECRET", "autocare_secret_key"),
		AccessTokenMinutes: accessMinutes,
		RefreshTokenDays:   refreshDays,
	}

	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	return cfg
}

func getEnv(key, fallback string) string {

	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}

func getEnvAsInt(name string, defaultValue int) int {

	valueStr := os.Getenv(name)

	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}

	return value
}
