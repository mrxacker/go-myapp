package config

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment string
	HTTPPort    int
	GRPCPort    int
	DBHost      string
	DBPort      int
	DBUser      string
	DBPassword  string
	DBName      string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("error loading .env file")
	}

	return &Config{
		Environment: getEnv("ENVIRONMENT", "development"),
		HTTPPort:    getEnvAsInt("HTTP_PORT", 8080),
		GRPCPort:    getEnvAsInt("GRPC_PORT", 9090),
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnvAsInt("DB_PORT", 5432),
		DBUser:      getEnv("DB_USER", "postgres"),
		DBPassword:  getEnv("DB_PASSWORD", ""),
		DBName:      getEnv("DB_NAME", "myapp"),
	}, nil
}

func ConnectDB(cfg *Config) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("database connection error: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		if i, err := strconv.Atoi(value); err == nil {
			return i
		}
	}
	return fallback
}
