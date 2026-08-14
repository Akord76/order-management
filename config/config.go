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
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		ServerPort: os.Getenv("SERVER_PORT"),

		JWTSecret:     os.Getenv("JWT_SECRET"),
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

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return fallback
}
