package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db DbConfig
}

type DbConfig struct {
	Dsn string
}

func LoadConfig() *Config {
	// прочитать .env
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file, using default config")
	}
	// Создаем зависимость
	return &Config{
		Db: DbConfig{
			Dsn: os.Getenv("DSN"),
		},
	}
}
