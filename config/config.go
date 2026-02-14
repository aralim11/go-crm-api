package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Version     string
	ServiceName string
	Database    DatabaseConfig
	Server      ServerConfig
	JWT         JWTConfig
}

type DatabaseConfig struct {
	Host     string
	DBPort   string
	User     string
	Password string
	DBName   string
}

type ServerConfig struct {
	AppUrl   string
	HTTPPort string
	Mode     string
}

type JWTConfig struct {
	Secret          string
	AccessTokenTTL  int
	RefreshTokenTTL int
}

func LoadConfig() (*Config, error) {
	godotenv.Load()

	config := &Config{
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			DBPort:   getEnv("DB_PORT", "5432"),
			User:     getEnv("DB_USER", "library_user"),
			Password: getEnv("DB_PASSWORD", ""),
			DBName:   getEnv("DB_NAME", "online_library"),
		},

		Server: ServerConfig{
			AppUrl:   getEnv("APP_URL", "localhost"),
			HTTPPort: getEnv("HTTP_PORT", "8081"),
			Mode:     getEnv("MODE", "debug"),
		},

		JWT: JWTConfig{
			Secret:          getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
			AccessTokenTTL:  24,
			RefreshTokenTTL: 168,
		},
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
