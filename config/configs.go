package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Configs struct {
	PostgresHost     string
	PostgresPort     string
	PostgresDB       string
	PostgresUsername string
	PostgresPassword string
}

var (
	config     *Configs
	configOnce sync.Once
)

func GetValue() *Configs {
	configOnce.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: .env file not found, using environment variables")
		}
		config = &Configs{
			PostgresHost:     GetENV("POSTGRES_HOST", "localhost"),
			PostgresPort:     GetENV("POSTGRES_PORT", "5432"),
			PostgresDB:       GetENV("POSTGRES_DB", "postgres"),
			PostgresUsername: GetENV("POSTGRES_USERNAME", "postgres"),
			PostgresPassword: GetENV("POSTGRES_PASSWORD", "password"),
		}
	})
	log.Default().Println("Config & ENV loaded successfully")
	return config
}

func GetENV(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
