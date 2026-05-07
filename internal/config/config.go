package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
	Port string
	DBUrl string
}


func LoadConfig() *Config{
	if err := godotenv.Load(); err != nil{
		log.Println("No .env file found, relying on system environment variables")
	}

	cfg := &Config{
		Port: os.Getenv("PORT"),
		DBUrl: getDBUrl(),
	}

	return cfg
}


func getEnv(key, defaultValue string) string{
	if value, exists := os.LookupEnv(key); exists{
		return value
	}
	return defaultValue
}



func getDBUrl() string {
    if url := os.Getenv("DATABASE_URL"); url != "" {
        return url
    }
    
	// Fallback to constructing it
	return "host=" + getEnv("DB_HOST", "localhost") +
		" user=" + getEnv("DB_USER", "postgres") +
		" password=" + getEnv("DB_PASSWORD", "postgres") +
		" dbname=" + getEnv("DB_NAME", "caronago") +
		" port=" + getEnv("DB_PORT", "5432") +
		" sslmode=disable"
}