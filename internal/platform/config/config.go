package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

type Config struct {
	HTTPAddr    string
	PostgresDSN string
	RedisAddr   string
}

func Load() *Config {
	return &Config{
		HTTPAddr:    os.Getenv("HTTPAddr"),
		PostgresDSN: os.Getenv("PostgresDSN"),
		RedisAddr:   os.Getenv("RedisAddr"),
	}
}
