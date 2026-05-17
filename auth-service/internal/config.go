package internal

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port               string
	DatabaseURL        string
	JWTSecret          string
	AccessTokenMinutes int
	RefreshTokenDays   int
}

func LoadConfig() Config {
	_ = godotenv.Load()

	accessMinutes, _ := strconv.Atoi(getEnv("ACCESS_TOKEN_MINUTES", "15"))
	refreshDays, _ := strconv.Atoi(getEnv("REFRESH_TOKEN_DAYS", "30"))

	return Config{
		Port:               getEnv("PORT", "8082"),
		DatabaseURL:        getEnv("DATABASE_URL", ""),
		JWTSecret:          getEnv("JWT_SECRET", "autocare_secret_key"),
		AccessTokenMinutes: accessMinutes,
		RefreshTokenDays:   refreshDays,
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
