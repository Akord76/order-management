package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// AppConfig holds every configuration value the application needs.
type AppConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	ServerPort string

	JWTSecret     string
	JWTExpireHour int
}

var Cfg *AppConfig

// LoadConfig reads a .env file (if present) and populates Cfg.
// It falls back to sane defaults when a variable is not set.
func LoadConfig() *AppConfig {
	if err := godotenv.Load(); err != nil {
		log.Println("config: no .env file found, relying on OS environment variables")
	}

	Cfg = &AppConfig{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     mustGetEnv("DB_PORT"),
		DBUser:     mustGetEnv("DB_USER"),
		DBPassword: mustGetEnv("DB_PASSWORD"),
		DBName:     getEnv("DB_NAME", "OrderManagement"),
		ServerPort: getEnv("SERVER_PORT", "8083"),

		JWTSecret:     mustGetEnv("JWT_SECRET"),
		JWTExpireHour: getEnvInt("JWT_EXPIRE_HOURS", 60),
	}

	return Cfg
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return fallback
}

// mustGetEnv is for secrets (JWT_SECRET, DB_PASSWORD) that must never have
// a hardcoded fallback baked into source code. If it's missing, fail fast
// at startup instead of silently running with an empty/known secret.
func mustGetEnv(key string) string {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		log.Fatalf("config: required env var %s is not set (check your .env file)", key)
	}
	return v
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return fallback
}
