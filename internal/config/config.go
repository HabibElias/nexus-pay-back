package config

import "os"

type Config struct {
	AppName    string
	HTTPPort   string
	GRPCPort   string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() *Config {
	LoadEnv() // optional, loads .env

	return &Config{
		AppName:    getEnv("APP_NAME", "TodoApp"),
		HTTPPort:   getEnv("HTTP_PORT", "8080"),
		GRPCPort:   getEnv("GRPC_PORT", "50052"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "nexusdb"),
	}
}

// getEnv reads an environment variable or returns default
func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
